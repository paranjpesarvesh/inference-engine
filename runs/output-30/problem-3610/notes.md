# Minimum Number of Primes to Sum to Target

## Description

Given an integer array nums and a target integer sum, find the minimum number of primes that add up to the target.

## Approaches

• Brute force: Check every element in the array and see if it is prime. If it is, add it to a set and check if the sum of all elements in the set equals the target.
• Optimized idea: Use a hash table to keep track of the sum of elements seen so far. When the current sum is equal to the target, count the number of unique elements in the hash table.

## Explanation

Brute force approach involves iterating through every element in the array and checking if it is prime. If it is, we add it to a set and check if the sum of all elements in the set equals the target. This approach has a time complexity of O(n^2) as we are iterating through each element twice.
The optimized idea involves using a hash table to keep track of the sum of elements seen so far. We start with an empty hash table and add the first element in the array to it. Then, we iterate through the rest of the array and for each element, we calculate its sum by adding it to the current sum stored in the hash table. If the current sum is equal to the target, we count the number of unique elements in the hash table. This approach has a time complexity of O(n) as we are iterating through the array once.
In both approaches, we need to check if an element is prime. We can use a simple algorithm like trial division to check for primality.

## Time Complexity

O(n^2) for brute force and O(n) for optimized idea.

## Flashcard

Minimum number of primes to sum to target: Use a hash table to keep track of the sum of elements seen so far.

---

Difficulty: 2 • Paid: yes