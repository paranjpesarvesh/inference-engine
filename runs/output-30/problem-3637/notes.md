# Trionic Array I

## Description

Given an array of integers, find the sum of all elements in it.

## Approaches

• Brute force: Iterate through each element and add it to a running total.
• Optimized idea: Use the prefix sum technique to calculate the sum of subarrays.
• Key tradeoffs: Brute force has O(n) time complexity, while optimized idea has O(n) time complexity.

## Explanation

The prefix sum technique involves iterating through the array and keeping track of the sum of subarrays. This allows us to calculate the sum of the entire array in constant time.

## Time Complexity

O(n)

## Flashcard

Prefix sum: Calculate the sum of subarrays to find the sum of the entire array.

---

Difficulty: 1