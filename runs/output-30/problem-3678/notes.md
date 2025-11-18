# Smallest Absent Positive Greater Than Average

## Description

Given an array of positive integers, find the smallest positive integer that is greater than the average of the given array.

## Approaches

• Brute force: Iterate through all positive integers and check if they are greater than the average.
• Optimized idea: Calculate the sum of all positive integers in the array and divide by the number of elements. Then, iterate through the array and return the smallest positive integer that is greater than the calculated sum divided by the number of elements.

## Explanation

The optimal solution involves calculating the sum of all positive integers in the array and dividing it by the number of elements. This gives us an average value for the array. We then iterate through the array and return the smallest positive integer that is greater than this calculated average value.

## Time Complexity

O(n)

## Flashcard

Average + Iterate

---

Difficulty: 1