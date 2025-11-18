# Find Zombie Sessions

## Description

Given a list of sessions and their start and end times, find all the zombie sessions where the last session ends before the next one starts.

## Approaches

• Brute force: Iterate through each session and check if it overlaps with any previous session. If there is an overlap, add the current session to the list of zombies.
• Optimized idea: Use a hash table to keep track of the end times of all sessions seen so far. When iterating through the sessions, check if the current session's start time is before the end time of any previous session in the hash table. If it is, add the current session to the list of zombies.
• Key tradeoff: The optimized idea has a time complexity of O(n log n) due to the use of a hash table. However, it requires additional memory to store the hash table.

## Explanation

First, we iterate through each session in the list.
For each session, we check if its start time is before the end time of any previous session in the list. If it is, we add the current session to the list of zombies.
We continue this process until we have checked all sessions in the list.

## Time Complexity

O(n log n)

## Flashcard

Zombie Sessions: Find sessions where last session ends before next one starts. Time complexity: O(n log n)

---

Difficulty: 3