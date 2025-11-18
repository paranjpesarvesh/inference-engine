# Maximum Balanced Shipments

## Description

Given a list of items with their respective weights and values, find the maximum possible shipment value that can be balanced in at most two boxes.

## Approaches

• Brute force: Try all possible combinations of items until a valid solution is found.
• Dynamic programming: Use memoization to store previously computed results and avoid redundant calculations.

## Explanation

The optimal solution involves sorting the items by weight and then iterating through them. For each item, we add it to the box with the smallest total weight so far. If the total weight exceeds the maximum allowed, we start a new box. We repeat this process until all items have been added to either box. The final value of the shipment is the sum of the values in both boxes.

## Time Complexity

O(n log n) for dynamic programming approach.

## Flashcard

Maximize shipment value by sorting items by weight and iterating through them.

---

Difficulty: 2