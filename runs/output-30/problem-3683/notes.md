# Earliest Time to Finish One Task

## Description

Given a list of tasks with their respective deadlines, find the earliest possible time to finish all tasks.

## Approaches

• Brute force: Sort tasks by deadline and iterate over them in order.
• Optimized idea: Use a priority queue to keep track of the latest task deadline.

## Explanation

The optimal solution is to use a priority queue. We can initialize an empty priority queue and iterate over the tasks. For each task, we calculate its earliest finish time by subtracting its duration from its deadline. If the earliest finish time is less than the current earliest finish time in the priority queue, we update it. Finally, we return the earliest finish time found in the priority queue.

## Time Complexity

O(n log n) for sorting and priority queue operations.

## Flashcard

Earliest Time to Finish One Task: Use a priority queue to keep track of tasks' deadlines.

---

Difficulty: 1