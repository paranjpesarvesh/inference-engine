# Sort Array By Absolute Value

## Description

Given an array of integers, sort it in ascending order based on the absolute value of each element.

## Approaches

• Brute force: Iterate through the array and compare each element with every other element. If two elements have different absolute values, swap them.
• Optimized idea: Use a priority queue to store the elements in order of their absolute values. Then iterate through the priority queue and add each element to the sorted array.

## Explanation

The optimal solution is to use a priority queue. We can create an empty priority queue, then iterate through the input array. For each element, we calculate its absolute value and compare it with the smallest absolute value in the priority queue. If our current element has a smaller absolute value, we remove the smallest element from the priority queue and add our current element to the front of the priority queue. This ensures that the smallest absolute values are always at the front of the priority queue. Once all elements have been added to the priority queue, we can iterate through it and add each element to the sorted array.

## Time Complexity

O(n log n) for both brute force and optimized idea.

## Flashcard

Sort an array by absolute value using a priority queue.

---

Difficulty: 1 • Paid: yes