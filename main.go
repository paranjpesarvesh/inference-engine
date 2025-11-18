package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxRetries   = 3
	baseDelaySec = 1
	failedFile   = "failed.json"
)

var (
	inputFile = flag.String("input", "problems.json", "Input JSON file")
	outDir    = flag.String("outdir", "output", "Output directory")
	textModel = flag.String("text-model", "mistral:7b-instruct-q4_K_M", "Text model name")
	codeModel = flag.String("code-model", "deepseek-coder:6.7b", "Code model name")
	workers   = flag.Int("workers", 4, "Concurrent prompt workers")
	ollamaURL = flag.String("ollama-url", "http://localhost:11434/api/generate", "Ollama API URL")
	logLevel  = flag.String("log-level", "info", "log level")
	testID    = flag.Int("test-id", -1, "Run in single-problem test mode for this ID")
)

var logger *slog.Logger
var httpClient = &http.Client{Timeout: 300 * time.Second}
var failedMu sync.Mutex
var failedCounter int32

type Problem struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Difficulty int    `json:"difficulty"`
	PaidOnly   bool   `json:"paid_only"`
	IsFavor    bool   `json:"is_favor"`
}

type Job struct{ Problem Problem }

type Result struct {
	Problem Problem
	Files   map[string]string
	Err     error
}

type InferenceRequest struct {
	Model  string
	Prompt string
	RespCh chan InferenceResponse
	Ctx    context.Context
}

type InferenceResponse struct {
	Text string
	Err  error
}

type simpleCache struct {
	mu sync.RWMutex
	m  map[string]string
}

func newCache() *simpleCache { return &simpleCache{m: make(map[string]string)} }
func (c *simpleCache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[k]
	return v, ok
}
func (c *simpleCache) Set(k, v string) {
	c.mu.Lock()
	c.m[k] = v
	c.mu.Unlock()
}

func main() {
	flag.Parse()
	lvl := slog.LevelInfo
	switch strings.ToLower(*logLevel) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	}
	logFile, err := os.OpenFile("pipeline.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		panic(fmt.Errorf("cannot open pipeline.log: %w", err))
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logger = slog.New(slog.NewJSONHandler(mw, &slog.HandlerOptions{Level: lvl}))
	slog.SetDefault(logger)

	logger.Info("pipeline starting", "text_model", *textModel, "code_model", *codeModel, "workers", *workers)

	runtime.GOMAXPROCS(runtime.NumCPU())
	if *testID != -1 {
		runSingleTestMode(*testID)
		return
	}

	problems, err := loadProblems(*inputFile)
	if err != nil {
		logger.Error("load problems failed", "error", err)
		return
	}
	if len(problems) == 0 {
		logger.Error("no problems found")
		return
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		logger.Error("mkdir failed", "error", err)
		return
	}

	jobs := make(chan Job, len(problems))
	results := make(chan Result, len(problems))

	textQ := make(chan InferenceRequest, 200)
	codeQ := make(chan InferenceRequest, 200)

	cache := newCache()

	var infWG sync.WaitGroup
	infWG.Add(2)
	go inferenceWorker(textQ, &infWG)
	go inferenceWorker(codeQ, &infWG)

	var wg sync.WaitGroup
	for i := 0; i < *workers; i++ {
		id := i
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			logger.Info("worker started", "worker_id", workerID)
			for job := range jobs {
				logger.Info("worker received job", "worker_id", workerID, "problem_id", job.Problem.ID)
				res := processJobWithRetries(job.Problem, textQ, codeQ, cache, workerID)
				results <- res
			}
			logger.Info("worker stopped", "worker_id", workerID)
		}(id)
	}

	for _, p := range problems {
		jobs <- Job{Problem: p}
	}
	close(jobs)

	var writerWG sync.WaitGroup
	writerWG.Add(1)
	go func() {
		defer writerWG.Done()
		for r := range results {
			if r.Err != nil {
				logger.Error("job failed", "id", r.Problem.ID, "error", r.Err)
				continue
			}
			if err := writeFiles(*outDir, r.Problem, r.Files); err != nil {
				logger.Error("writeFiles failed", "id", r.Problem.ID, "error", err)
			}
		}
	}()

	wg.Wait()
	close(results)
	writerWG.Wait()

	close(textQ)
	close(codeQ)
	infWG.Wait()

	logger.Info("pipeline completed")
}

// processJobWithRetries wraps attempts with exponential backoff, partial saves, and failure registry.
func processJobWithRetries(p Problem, textQ, codeQ chan InferenceRequest, cache *simpleCache, workerID int) Result {
	start := time.Now()
	logger.Info("process start (with retries)", "id", p.ID, "worker", workerID)

	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		logger.Info("attempt", "id", p.ID, "attempt", attempt)
		textOut, codeOut, err := processJobAttempt(p, textQ, codeQ, cache, workerID)
		if err == nil {
			files := buildFilesFromResponses(p, textOut, codeOut)
			logger.Info("process success", "id", p.ID, "worker", workerID, "elapsed_ms", time.Since(start).Milliseconds())
			return Result{Problem: p, Files: files}
		}

		lastErr = err
		logger.Warn("attempt failed", "id", p.ID, "attempt", attempt, "error", err)

		if textOutIsUseful(textOut) && !codeOutIsUseful(codeOut) {
			partial := buildFilesFromResponses(p, textOut, "") // code empty
			if err := writeFiles(*outDir, p, partial); err != nil {
				logger.Warn("partial write failed", "id", p.ID, "err", err)
			} else {
				dir := filepath.Join(*outDir, fmt.Sprintf("problem-%d", p.ID))
				os.WriteFile(filepath.Join(dir, "PARTIAL_NOTE_ONLY"), []byte(time.Now().Format(time.RFC3339)), 0o644)
				logger.Info("wrote partial notes.md", "id", p.ID)
			}
		}

		if attempt < maxRetries {
			sleep := time.Duration(math.Pow(2, float64(attempt-1))*baseDelaySec) * time.Second
			logger.Info("retry backoff", "id", p.ID, "sleep_s", sleep.Seconds())
			time.Sleep(sleep)
		}
	}

	recordFailure(p, lastErr)

	logger.Error("all attempts failed", "id", p.ID, "error", lastErr)
	return Result{Problem: p, Err: fmt.Errorf("all attempts failed: %w", lastErr)}
}

// processJobAttempt does a single try (no retries). Returns textOut, codeOut, or error.
func processJobAttempt(p Problem, textQ, codeQ chan InferenceRequest, cache *simpleCache, workerID int) (string, string, error) {
	start := time.Now()
	logger.Debug("processJobAttempt start", "id", p.ID, "worker", workerID)

	keyText := fmt.Sprintf("text-%d", p.ID)
	keyCode := fmt.Sprintf("code-%d", p.ID)

	var textOut string
	if v, ok := cache.Get(keyText); ok {
		textOut = v
	} else {
		prompt := buildTextPromptStrict(p)
		out, err := callInference(textQ, *textModel, prompt)
		if err != nil {
			return "", "", fmt.Errorf("text inference error: %w", err)
		}

		if s, err := enforceJSON(textSchema, out); err == nil {
			textOut = s
		} else {
			logger.Warn("text JSON invalid, storing fallback", "id", p.ID, "err", err)
			textOut = fallbackDescription(out)
		}
		cache.Set(keyText, textOut)
	}

	var codeOut string
	if v, ok := cache.Get(keyCode); ok {
		codeOut = v
	} else {
		shortSummary := summarizeForCode(textOut)
		prompt := buildCodePromptStrict(p, shortSummary)
		out, err := callInference(codeQ, *codeModel, prompt)
		if err != nil {
			return textOut, "", fmt.Errorf("code inference error: %w", err)
		}

		codeOut = sanitizeCPP(out)
		if !isBalancedBraces(codeOut) {
			logger.Warn("code unbalanced braces, applying auto-repair (id)", "id", p.ID)
			codeOut = repairBraces(codeOut)
		}
		if !codeOutIsUseful(codeOut) {
			return textOut, codeOut, fmt.Errorf("code validation failed")
		}
		cache.Set(keyCode, codeOut)
	}

	logger.Debug("processJobAttempt done", "id", p.ID, "elapsed_ms", time.Since(start).Milliseconds())
	return textOut, codeOut, nil
}

// textOutIsUseful checks if the text output is worth saving as partial notes (basic heuristics).
func textOutIsUseful(text string) bool {
	if strings.TrimSpace(text) == "" {
		return false
	}
	if _, err := enforceJSON(textSchema, text); err == nil {
		return true
	}
	return len(strings.TrimSpace(text)) > 100
}

func codeOutIsUseful(code string) bool {
	s := strings.TrimSpace(code)
	if s == "" {
		return false
	}
	if !strings.Contains(s, "#include") {
		return false
	}
	if !strings.Contains(s, "class Solution") {
		return false
	}
	if !isBalancedBraces(s) {
		return false
	}
	if len(s) < 80 {
		return false
	}
	return true
}

// summarizeForCode produces a short summary (one-liner) from text JSON to reduce prompt size.
func summarizeForCode(text string) string {
	if s, err := enforceJSON(textSchema, text); err == nil {
		var doc map[string]any
		_ = json.Unmarshal([]byte(s), &doc)
		desc := ""
		if v, ok := doc["description"]; ok {
			desc = toString(v)
			if len(desc) > 300 {
				desc = desc[:300]
			}
		}
		return strings.ReplaceAll(strings.Join([]string{desc}, " "), "\n", " ")
	}
	if len(text) > 300 {
		return text[:300]
	}
	return text
}

// recordFailure appends failure metadata to runs/v{version}/failed.json (creates file if needed)
func recordFailure(p Problem, err error) {
	failedMu.Lock()
	defer failedMu.Unlock()

	base := *outDir
	if base == "" {
		base = "output"
	}
	if err := os.MkdirAll(base, 0o755); err != nil {
		logger.Warn("recordFailure mkdir failed", "err", err)
	}

	path := filepath.Join(base, failedFile)
	var arr []map[string]any
	if b, e := os.ReadFile(path); e == nil {
		_ = json.Unmarshal(b, &arr)
	}

	entry := map[string]any{
		"id":     p.ID,
		"title":  p.Title,
		"reason": fmt.Sprint(err),
		"time":   time.Now().Format(time.RFC3339),
	}
	arr = append(arr, entry)

	nb, _ := json.MarshalIndent(arr, "", "  ")
	if e := os.WriteFile(path, nb, 0o644); e != nil {
		logger.Warn("failed to write failed.json", "err", e)
	} else {
		atomic.AddInt32(&failedCounter, 1)
		logger.Info("recorded failure", "id", p.ID)
	}
}

// Strict prompts
var textSchema = []string{"description", "approaches", "explanation", "time_complexity", "flashcard"}

func buildTextPromptStrict(p Problem) string {
	return fmt.Sprintf(`You MUST output a single valid JSON object with these keys.
		You MUST output a single valid JSON object with these keys:

		{
		  "description": "2–4 sentences explaining the problem clearly.",
		  "approaches": "Bullet list (• or -) covering brute force, optimized ideas, and key tradeoffs.",
		  "explanation": "Step-by-step reasoning of the optimal solution: intuition, important observations, corner cases.",
		  "time_complexity": "Big-O time and space with a short justification.",
		  "flashcard": "1–2 line memory hook with the core idea."
		}

		Rules:
		- Output ONLY the JSON (no markdown, no backticks, no extra text).
		- Be clear, complete, and educational.
		- Provide strong reasoning, not superficial statements.
		- Bullets allowed only in 'approaches'.

		Problem to solve %d: %s`, p.ID, p.Title)
}

func buildCodePromptStrict(p Problem, summary string) string {
	return fmt.Sprintf(`Generate ONLY valid, compilable C++17 code for LeetCode. Output must be plain code text (no markdown or comments).
		You are an expert competitive programmer.
		Produce ONLY valid, compilable C++17 code.
		STRICT RULES (must obey):
		- No markdown
		- No backticks
		- No comments
		- No explanation text
		- Must use the correct LeetCode method signature.

		Fill this structure exactly:

		#include <bits/stdc++.h>
		using namespace std;

		class Solution {
		public:
			<METHOD_SIGNATURE> {
				<CODE>
			}
		};

		Context summary (from assistant): %s
		Problem: %s`, summary, p.Title)
}

// Inference helpers
func callInference(q chan InferenceRequest, model, prompt string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Second)
	defer cancel()

	req := InferenceRequest{Model: model, Prompt: prompt, RespCh: make(chan InferenceResponse, 1), Ctx: ctx}

	select {
	case q <- req:
	case <-ctx.Done():
		return "", ctx.Err()
	}

	select {
	case resp := <-req.RespCh:
		return resp.Text, resp.Err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func inferenceWorker(q chan InferenceRequest, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range q {
		out, err := inferOnce(req.Ctx, req.Model, req.Prompt)
		req.RespCh <- InferenceResponse{Text: out, Err: err}
	}
}

func inferOnce(ctx context.Context, model, prompt string) (string, error) {
	np := optimalNumPredict(model, prompt)

	payload := map[string]any{
		"model":  model,
		"prompt": prompt,
		"stream": true,
		"options": map[string]any{
			"temperature": 0.2,
			"num_predict": np,
			"top_p":       0.95,
		},
	}
	b, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, "POST", *ollamaURL, bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama HTTP %d: %s", resp.StatusCode, string(body))
	}

	dec := json.NewDecoder(resp.Body)
	var combined strings.Builder

	for {
		var obj map[string]any
		if err := dec.Decode(&obj); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", fmt.Errorf("decode error: %w", err)
		}

		if v, ok := obj["response"]; ok {
			combined.WriteString(fmt.Sprint(v))
		}

		if v, ok := obj["delta"]; ok {
			combined.WriteString(fmt.Sprint(v))
		}

		if done, ok := obj["done"].(bool); ok && done {
			break
		}
	}

	res := strings.TrimSpace(combined.String())
	if res == "" {
		return "", fmt.Errorf("empty combined response")
	}

	logger.Debug("collected response", "model", model, "len", len(res), "preview", preview(res, 200))
	return res, nil
}

// utilities
func optimalNumPredict(model, prompt string) int {
	lm := strings.ToLower(model)
	plen := len(prompt)
	if strings.Contains(lm, "phi3") || strings.Contains(lm, "phi") {
		if plen < 1200 {
			return 550
		}
		return 650
	}
	if strings.Contains(lm, "deepseek") || strings.Contains(lm, "coder") || strings.Contains(lm, "code") {
		if plen < 1500 {
			return 450
		}
		return 600
	}
	if strings.Contains(lm, "mistral") {
		return 650
	}
	return 650
}

// enforceJSON: extract JSON substring and validate required keys
func enforceJSON(schema []string, raw string) (string, error) {
	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")
	if start == -1 || end == -1 || end <= start {
		return "", fmt.Errorf("no JSON object found")
	}
	jsonStr := raw[start : end+1]
	var doc map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &doc); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}
	for _, k := range schema {
		if _, ok := doc[k]; !ok {
			return "", fmt.Errorf("missing field: %s", k)
		}
	}
	b, _ := json.Marshal(doc)
	return string(b), nil
}

var reLineComment = regexp.MustCompile(`(?m)//.*$`)
var reBlockComment = regexp.MustCompile(`(?s)/\*.*?\*/`)

func sanitizeCPP(raw string) string {
	r := strings.TrimSpace(raw)
	r = strings.ReplaceAll(r, "```cpp", "")
	r = strings.ReplaceAll(r, "```", "")
	r = reBlockComment.ReplaceAllString(r, "")
	r = reLineComment.ReplaceAllString(r, "")
	r = strings.TrimSpace(r)
	lines := strings.Split(r, "\n")
	clean := make([]string, 0, len(lines))
	blank := false
	for _, ln := range lines {
		if strings.TrimSpace(ln) == "" {
			if !blank {
				clean = append(clean, "")
			}
			blank = true
			continue
		}
		blank = false
		clean = append(clean, ln)
	}
	return strings.Join(clean, "\n")
}

func isBalancedBraces(s string) bool {
	count := 0
	for _, ch := range s {
		switch ch {
		case '{':
			count++
		case '}':
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

func repairBraces(s string) string {
	count := 0
	for _, ch := range s {
		switch ch {
		case '{':
			count++
		case '}':
			count--
		}
	}
	if count > 0 {
		for i := 0; i < count; i++ {
			s += "\n}"
		}
	}
	return s
}

func fallbackDescription(raw string) string {
	r := strings.TrimSpace(raw)
	r = strings.ReplaceAll(r, "\r\n", "\n")
	r = strings.ReplaceAll(r, "\r", "\n")
	parts := strings.SplitN(r, "\n\n", 2)
	if len(parts) > 0 && strings.TrimSpace(parts[0]) != "" {
		p := strings.TrimSpace(parts[0])
		if len(p) > 800 {
			return p[:800]
		}
		return p
	}
	if len(r) > 800 {
		return r[:800]
	}
	return r
}

func writeFiles(base string, p Problem, files map[string]string) error {
	dir := filepath.Join(base, fmt.Sprintf("problem-%d", p.ID))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	normalize := func(s string) string {
		if s == "" {
			return s
		}
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, "\r\n", "\n")
		s = strings.ReplaceAll(s, "\r", "\n")
		lines := strings.Split(s, "\n")
		cleaned := make([]string, 0, len(lines))
		blank := false
		for _, ln := range lines {
			if strings.TrimSpace(ln) == "" {
				if !blank {
					cleaned = append(cleaned, "")
				}
				blank = true
				continue
			}
			blank = false
			cleaned = append(cleaned, ln)
		}
		return strings.Join(cleaned, "\n")
	}

	for name, content := range files {
		content = normalize(content)
		if strings.TrimSpace(content) == "" || content == "null" {
			logger.Debug("skipping empty file", "id", p.ID, "file", name)
			continue
		}
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return err
		}
		logger.Info("wrote file", "id", p.ID, "file", name, "bytes", len(content))
	}
	return nil
}

func preview(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func buildFilesFromResponses(p Problem, textResp, codeResp string) map[string]string {
	files := make(map[string]string)

	var doc map[string]any
	if err := json.Unmarshal([]byte(textResp), &doc); err == nil {
		notes := buildUnifiedNotes(doc, p)
		files["notes.md"] = notes
	} else {
		fb := fallbackDescription(textResp)
		notes := fmt.Sprintf("# %s\n\n## Description\n\n%s\n", p.Title, normalizeMarkdown(fb))
		files["notes.md"] = normalizeMarkdown(notes)
	}

	files["solution.cpp"] = codeResp

	return files
}

// buildUnifiedNotes transforms the strict JSON object into a polished notes.md
func buildUnifiedNotes(doc map[string]any, p Problem) string {
	title := p.Title
	var sb strings.Builder

	sb.WriteString("# " + title + "\n\n")

	if v, ok := doc["description"]; ok {
		desc := toString(v)
		sb.WriteString("## Description\n\n")
		sb.WriteString(normalizeMarkdown(desc) + "\n\n")
	}

	if v, ok := doc["approaches"]; ok {
		app := toString(v)
		sb.WriteString("## Approaches\n\n")
		sb.WriteString(normalizeBullets(app) + "\n\n")
	}

	if v, ok := doc["explanation"]; ok {
		exp := toString(v)
		sb.WriteString("## Explanation\n\n")
		sb.WriteString(normalizeMarkdown(exp) + "\n\n")
	}

	if v, ok := doc["time_complexity"]; ok {
		tc := toString(v)
		sb.WriteString("## Time Complexity\n\n")
		sb.WriteString(normalizeMarkdown(tc) + "\n\n")
	}

	if v, ok := doc["flashcard"]; ok {
		fc := toString(v)
		sb.WriteString("## Flashcard\n\n")
		sb.WriteString(strings.TrimSpace(fc) + "\n\n")
	}

	meta := []string{}
	if p.Difficulty > 0 {
		meta = append(meta, fmt.Sprintf("Difficulty: %d", p.Difficulty))
	}
	if p.PaidOnly {
		meta = append(meta, "Paid: yes")
	}
	if p.IsFavor {
		meta = append(meta, "Favorite: yes")
	}
	if len(meta) > 0 {
		sb.WriteString("---\n\n")
		sb.WriteString(strings.Join(meta, " • ") + "\n")
	}

	return normalizeMarkdown(sb.String())
}

// toString safely converts JSON values to strings
func toString(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case []any:
		parts := make([]string, 0, len(t))
		for _, el := range t {
			parts = append(parts, toString(el))
		}
		return strings.Join(parts, "\n")
	case map[string]any:
		b, _ := json.Marshal(t)
		return string(b)
	default:
		return fmt.Sprint(t)
	}
}

// normalizeBullets ensures approaches appear as neat bullet list with '•' bullets
func normalizeBullets(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	lines := strings.Split(s, "\n")
	out := make([]string, 0, len(lines))
	for _, ln := range lines {
		ln = strings.TrimSpace(ln)
		if ln == "" {
			if len(out) > 0 && out[len(out)-1] != "" {
				out = append(out, "")
			}
			continue
		}
		ln = strings.TrimLeft(ln, "-*• \t")
		ln = strings.TrimSpace(ln)
		out = append(out, "• "+ln)
	}
	if len(out) > 0 && out[0] == "" {
		out = out[1:]
	}
	if len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}
	return strings.Join(out, "\n")
}

// normalizeMarkdown collapses excessive blank lines and normalizes newlines
func normalizeMarkdown(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	re := regexp.MustCompile(`\n{3,}`)
	s = re.ReplaceAllString(s, "\n\n")
	return strings.TrimSpace(s) + "\n"
}

func loadProblems(path string) ([]Problem, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var ps []Problem
	if err := json.Unmarshal(b, &ps); err != nil {
		return nil, err
	}
	return ps, nil
}

func runSingleTestMode(id int) {
	logger.Info("TEST MODE ENABLED", "problem_id", id)

	problems, err := loadProblems(*inputFile)
	if err != nil {
		logger.Error("load problems failed", "err", err)
		return
	}

	var target *Problem
	for _, p := range problems {
		if p.ID == id {
			target = &p
			break
		}
	}

	if target == nil {
		logger.Error("test ID not found", "id", id)
		return
	}

	// Build prompts
	textPrompt := buildTextPromptStrict(*target)
	codePrompt := buildCodePromptStrict(*target, "This summary will be updated after text phase")

	logger.Info("running text inference test")
	text, err := inferOnce(context.Background(), *textModel, textPrompt)
	if err != nil {
		logger.Error("text inference failed", "error", err)
		return
	}

	logger.Info("text output", "preview", preview(text, 200))

	logger.Info("running code inference test")
	code, err := inferOnce(context.Background(), *codeModel, codePrompt)
	if err != nil {
		logger.Error("code inference failed", "error", err)
		return
	}

	logger.Info("code output", "preview", preview(code, 200))

	os.MkdirAll("test_output", 0o755)
	os.WriteFile(fmt.Sprintf("test_output/text_%d.json", id), []byte(text), 0o644)
	os.WriteFile(fmt.Sprintf("test_output/code_%d.cpp", id), []byte(code), 0o644)

	logger.Info("test mode completed")
}
