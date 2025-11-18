# Sum of Weighted Modes in Subarrays

## Description

Given an array of integers representing the weights and frequencies of items in a bag, find the sum of weighted modes in non-overlapping subarrays.

## Approaches

• Brute force: Iterate through all possible subarrays and calculate their sums.
• Optimized idea: Use dynamic programming to store the sum of each item in a subarray, then iterate through the array to find the mode.

## Explanation

Brute force approach involves iterating through all possible subarrays and calculating their sums. This can be done using nested loops, where the outer loop iterates over the start of the subarray and the inner loop iterates over the end of the subarray. The time complexity of this approach is O(n^3), where n is the length of the input array.
The optimized idea involves using dynamic programming to store the sum of each item in a subarray. We can create an array dp[] of size equal to the maximum possible sum of items in the bag, initialized to 0. Then, we iterate through the input array and for each element, we update the dp[] array with the maximum possible sum that can be obtained by including or excluding the current element. Finally, we iterate through the dp[] array to find the mode.
The time complexity of this optimized idea is O(n^2), where n is the length of the input array. This approach has a much better time complexity than the brute force approach.

## Time Complexity

O(n^2)

## Flashcard

Find the sum of weighted modes in non-overlapping subarrays using dynamic programming.

---

Difficulty: 2 • Paid: yes