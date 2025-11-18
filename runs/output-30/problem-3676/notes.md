# Count Bowl Subarrays

## Description

Given an array of integers representing the number of items in each bowl, count the number of subarrays that can be formed.

## Approaches

• Brute force: Iterate through all possible subarrays and count them.
• Optimized idea: Use dynamic programming to store counts of subarrays ending at each index.

## Explanation

The optimal solution involves using a sliding window approach. We start with an empty subarray and iterate over the array, adding items to the subarray until we exceed the maximum capacity. At each step, we update the count of subarrays ending at that index by including the current item or excluding it based on whether the subarray is full. This approach has a time complexity of O(n).

## Time Complexity

O(n)

## Flashcard

Count subarrays with dynamic programming.

---

Difficulty: 2