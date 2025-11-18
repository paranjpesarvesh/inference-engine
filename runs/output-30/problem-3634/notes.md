# Minimum Removals to Balance Array

## Description

Given an array of integers, find the minimum number of operations (insertions or deletions) required to balance it.

## Approaches

• Brute force: Try all possible insertions and deletions until the array is balanced.
• Optimized idea: Use two pointers to keep track of the sum of elements on the left and right sides. Insert or delete elements from the side with a smaller sum until both sides are equal.

## Explanation

Brute force approach involves trying all possible insertions and deletions on the array until it is balanced. This can be computationally expensive for large arrays.
The optimized idea uses two pointers to keep track of the sum of elements on the left and right sides. We start by iterating through the array from left to right. If the sum of elements on the left side is smaller than the sum on the right side, we insert an element on the left side. Otherwise, we delete an element on the right side. We continue this process until both sides have the same sum.
This approach has a time complexity of O(n) because we only need to iterate through the array once. The space complexity is also O(1) because we only need two pointers to keep track of the sums.

## Time Complexity

O(n)

## Flashcard

Balance an array with minimum operations using two pointers.

---

Difficulty: 2