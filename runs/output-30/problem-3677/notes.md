# Count Binary Palindromic Numbers

## Description

Count the number of binary palindromic numbers.

## Approaches

• Brute force: Check each number in the range for palindrome property.
• Optimized idea: Use a stack to check for palindrome property.

## Explanation

The brute force approach involves checking each number in the range for its palindrome property. This can be done by converting the number to string and comparing it with its reverse. However, this approach has a time complexity of O(n^2) where n is the range size. The optimized idea uses a stack to check for the palindrome property. We start from the least significant bit (LSB) and push it onto the stack. Then we compare the LSB with the most significant bit (MSB). If they are equal, we pop from the stack and move to the next bit. This process continues until we reach the middle of the number. The time complexity of this approach is O(n).

## Time Complexity

O(n)

## Flashcard

Count binary palindromic numbers using a stack.

---

Difficulty: 3