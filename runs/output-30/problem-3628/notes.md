# Maximum Number of Subsequences After One Inserting

## Description

Given an array of integers, find the maximum number of subsequences that can be obtained by inserting a single element into the array.

## Approaches

• Brute force: Generate all possible subarrays and check if they can be obtained by inserting a single element.
• Optimized idea: Use dynamic programming to store the maximum number of subsequences ending at each index, then iterate through the array and add the maximum of the current index and the previous index plus one.

## Explanation

The optimal solution involves iterating through the array and keeping track of the maximum number of subsequences that can end at each index. We can do this using dynamic programming, where we store the maximum number of subsequences ending at each index in an array. Then, we iterate through the input array and for each element, we add the maximum of the current index and the previous index plus one to our dynamic programming array. This gives us the maximum number of subsequences that can be obtained by inserting a single element into the array.

## Time Complexity

O(n^2)

## Flashcard

Maximum number of subsequences after inserting a single element.

---

Difficulty: 2