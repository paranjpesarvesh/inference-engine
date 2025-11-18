# Maximum Total from Optimal Activation Order

## Description

Given a list of n integers representing the activation cost of each machine, find the optimal order of activation that maximizes the total sum of activation costs.

## Approaches

• Brute force: Try all possible orders and calculate the total sum for each.
• Optimized idea: Use dynamic programming to build an array representing the minimum cost to activate machines up to a given index.

## Explanation

The optimal solution involves using dynamic programming to build an array representing the minimum cost to activate machines up to a given index. We can then iterate through this array and find the maximum sum of activation costs by adding the cost of each machine in the optimal order.

## Time Complexity

O(n^2)

## Flashcard

Minimum cost dynamic programming + iterative solution

---

Difficulty: 2