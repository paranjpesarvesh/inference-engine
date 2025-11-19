# Local LLM Automation Engine

This project is a fully local, concurrent pipeline that reads LeetCode problems, generates rich explanations using a text LLM, produces clean C++17 solutions using a code LLM, and writes deterministic output — all running on CPU.

It began as an attempt to solve every LeetCode problem in a single day, but it grew into a deeper exploration of:

- Local LLM inference
- Concurrency patterns in Go
- Streaming NDJSON decoding
- CPU bottlenecks
- Error-handling and retry systems
- Building a simple but effective local inference engine

Although the original challenge proved impossible, the resulting system is clean, fast, and surprisingly reliable.

---

# Repository Structure
```
.
├── resources/
├── runs/
│   └── output-30/                 (Sample output of 30 generated problems)
├── docker-compose.yaml             (Ollama runtime configuration)
├── go.mod
├── go.sum
├── main.go                         (Entire pipeline in one file)
├── pipeline.log                    (Sample logs)
├── problems.json                   (Primary problem list)
└── test30.json                     (Subset used for testing)
```
---

# What This Pipeline Does

For each LeetCode problem, the pipeline:

1. Loads the problem metadata
2. Constructs a strict JSON-based prompt
3. Calls a text model to generate:
   * Description
   * Approaches
   * Explanation
   * Time complexity
   * Flashcard summary
4. Validates the JSON
5. Summarizes the explanation
6. Calls a code model to produce a clean C++17 solution
7. Sanitizes and repairs the generated code
8. Writes two files for each problem:

runs/output-30/problem-<id>/
notes.md
solution.cpp

---

# Technical Highlights

**Single-file Go engine (<1000 LOC)**
The entire system is implemented in main.go for clarity and hackability.

**Dual inference queues (textQ and codeQ)**
CPU LLMs cannot parallelize inference effectively. Each queue has one dedicated inference worker.

**Concurrent prompt workers**
Goroutines handle preprocessing, JSON validation, postprocessing, and writing in parallel.

**Streaming NDJSON decoding**
Ollama responses are consumed incrementally for reduced latency.

**Strict JSON schema enforcement**
Explanations follow a predictable schema to ensure clean Markdown generation.

**Retry engine with exponential backoff**
Automatically handles malformed JSON, stream interruptions, and partial failures.

**C++ sanitizer and code-repair logic**
Ensures valid includes, balanced braces, and consistent formatting.

---

# Performance Results

On a CPU-only system with lightweight Ollama models:

- 30 problems processed in 1 hour
- Explanations: 100 percent correct
- Incorrect C++ solutions: 5
- Partial outputs: 1
- Average processing speed: roughly 2 minutes per problem
- Zero unrecoverable failures during the run

All of this was achieved without GPUs, batching, or parallel inference.

---

# The Real Bottleneck: CPU LLM Inference

Key observations from the project:

- One inference saturates all CPU cores
- Token generation is inherently sequential
- NUM_PARALLEL greater than 1 reduces throughput
- Prompt and output size heavily affect latency
- Even small models place significant load on CPUs
- Concurrency helps around inference but cannot speed up inference itself
- Modern LLM runtimes are GPU-first; dedicated CPU-first pipelines are still lacking

This pipeline effectively serves as a prototype for what a CPU-first LLM engine might look like.

---

# Running the Pipeline

Start Ollama:
```
docker compose up -d
```

Run the pipeline:
```
go run main.go -input problems.json -outdir runs/output-30
```

Run a single test problem:
```
go run main.go -test-id 1
```
---

# Configuration Flags

- -input: Path to problem list JSON
- -outdir: Output directory
- -test-id: Run one problem in isolation
- -workers: Number of concurrent workers (not inference workers)
- -log: Write pipeline logs to file

---

# Example Output

runs/output-30/problem-438/
notes.md
solution.cpp

---

# Example Problem Entry (problems.json)
```
  {
    "id": 2000,
    "title": "Reverse Prefix of Word",
    "slug": "reverse-prefix-of-word",
    "difficulty": 1,
    "paid_only": false,
    "is_favor": false
  }
```
---

# Why This Project Exists

This began as an attempt to automatically solve every LeetCode problem in a day.
Instead, it became a study in pipeline engineering, local inference, CPU constraints, and concurrency.

---

# Contribute

Suggestions, issues, and pull requests are welcome.
Experiment with alternative models, improve the sanitizer, or try adapting the pipeline for GPU inference.

