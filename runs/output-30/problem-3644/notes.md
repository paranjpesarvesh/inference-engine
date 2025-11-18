# Maximum K to Sort a Permutation

## Description

Given an array of integers and a positive integer k, find the maximum value of k that can be obtained by sorting the array.

## Approaches

• Brute force: Sort the array in ascending order and return the maximum value of k.
• Optimized idea: Use two pointers to keep track of the minimum and maximum values seen so far. Update the maximum value of k as we iterate through the array.

## Explanation

Brute force approach: Sorting the array in ascending order takes O(n log n) time complexity using any sorting algorithm. After sorting, we can return the maximum value of k.
Optimized idea approach: We use two pointers, one for minimum and one for maximum values seen so far. We initialize the minimum value to be the first element in the array and the maximum value to be the last element in the array. As we iterate through the array, we update the maximum value of k if it is greater than the current maximum value of k. This approach takes O(n) time complexity.

## Time Complexity

O(n log n) for brute force and O(n) for optimized idea.

## Flashcard

Find the maximum value of k by sorting an array in ascending order or using two pointers to keep track of minimum and maximum values seen so far.

---

Difficulty: 2