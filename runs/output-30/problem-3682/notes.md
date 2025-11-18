# Minimum Index Sum of Common Elements

## Description

Given an array of integers, find the minimum number of operations (insertions or deletions) required to transform it into a sorted array.

## Approaches

• Brute force: Iterate through all possible insertions and deletions and count the number of operations.
• Optimized idea: Use two pointers, one at the beginning and one at the end of the array, to keep track of the minimum sum of common elements.

## Explanation

Brute force approach involves iterating through all possible insertions and deletions. This can be done using nested loops, where the outer loop iterates over the elements in the array and the inner loop iterates over all possible positions to insert or delete an element.
The optimized idea uses two pointers, one at the beginning and one at the end of the array. The minimum sum of common elements is maintained by keeping track of the smallest element on the left side of the current pointer and the largest element on the right side of the current pointer. As we move the pointers towards each other, we can update the minimum sum of common elements by adding or subtracting the current element at the intersection of the two pointers.

## Time Complexity

O(n^2) for brute force and O(n) for optimized idea.

## Flashcard

Minimum Index Sum of Common Elements: Two pointers approach to maintain minimum sum of common elements.

---

Difficulty: 2 • Paid: yes