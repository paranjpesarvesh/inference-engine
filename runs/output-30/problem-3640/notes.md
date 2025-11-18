# Trionic Array II

## Description

Given an array of integers, find the sum of all elements that are not equal to their previous element.

## Approaches

• Brute force: Iterate through the array and compare each element with its previous element. If they are different, add them to a running total.
• Optimized idea: Use two pointers - one at the beginning of the array and one at the end. Compare the current element with the previous element and update the running total accordingly.

## Explanation

The optimal solution is to use two pointers, where we start from the beginning and end of the array. We compare the current element with the previous element and update the running total accordingly. This approach has a time complexity of O(n) because we only need to iterate through the array once.

## Time Complexity

O(n)

## Flashcard

Compare current element with previous element and update running total.

---

Difficulty: 3