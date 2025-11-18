# Number of Stable Subsequences

## Description

Given a string s, return the number of stable subsequences of s.

## Approaches

• Brute force: generate all substrings and check if they are stable.
• Optimized idea: use dynamic programming to count the number of stable suffixes of each prefix.

## Explanation

The optimal solution uses dynamic programming. We first count the number of stable suffixes of each prefix by iterating through the string and keeping track of the count for each substring. Then, we use these counts to count the number of stable subsequences of the entire string.

## Time Complexity

O(n^2)

## Flashcard

Count stable suffixes of each prefix to find number of stable subsequences.

---

Difficulty: 3