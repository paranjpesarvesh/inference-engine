#  Minimum Discards to Balance Inventory

## Description

Given an inventory of n discs with weights w1, w2, ..., wn where each disc can be picked up at most once and put down at any position on the stack. The goal is to find the minimum number of moves required to balance the stack.

## Approaches

• Brute force: Try all possible orders of picking up and putting down discs until a balanced stack is achieved.
• Optimized idea: Use dynamic programming to build an array dp[] such that dp[i] represents the minimum number of moves required to balance the stack after picking up disc i.

## Explanation

Step-by-step reasoning: Initialize an empty stack. For each disc i, calculate the minimum number of moves required to balance the stack after picking up disc i by considering two cases:
Case 1: The top disc on the stack has weight less than or equal to wi. In this case, we can simply pick up disc i and put it on top of the stack. The minimum number of moves required is 1 + dp[j-1] where j is the index of the top disc on the stack.
Case 2: The top disc on the stack has weight greater than wi. In this case, we need to remove the top disc from the stack and put it at the bottom. We can then pick up disc i and put it on top of the stack. The minimum number of moves required is 1 + dp[j-1] + dp[k-1] where j is the index of the top disc on the stack, k is the index of the second top disc on the stack, and dp[k-1] represents the minimum number of moves required to balance the stack after removing the second top disc.
Corner cases: If there are no discs left on the stack or if all discs have weight greater than wi, then the minimum number of moves required is infinity.

## Time Complexity

O(n^2)

## Flashcard

Minimum Discs to Balance Inventory: Dynamic Programming

---

Difficulty: 2