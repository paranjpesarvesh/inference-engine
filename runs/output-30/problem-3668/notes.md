# Restore Finishing Order

## Description

Given a list of numbers and their corresponding positions, restore the original finishing order of the numbers.

## Approaches

• Brute force: Sort the list of pairs in ascending order based on position.
• Optimized idea: Use a map to store the index of each number and sort the list based on the values in the map.

## Explanation

The optimal solution is to use a map to store the index of each number. We can then sort the list based on the values in the map, which will give us the original finishing order.

## Time Complexity

O(n log n) for both brute force and optimized idea.

## Flashcard

Map numbers to their positions, sort by values in map.

---

Difficulty: 1