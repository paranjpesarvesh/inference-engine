# Find Loyal Customers

## Description

Identify customers who consistently purchase products from a store.

## Approaches

• Brute force: Iterate through all customers and check their purchase history.
• Optimized idea: Use a hash table to keep track of customer purchases and count the number of unique items purchased.
• Key tradeoffs: Hash tables have O(1) average-case lookup time, but can be expensive to initialize for large datasets.

## Explanation

To find loyal customers, we need to identify those who consistently purchase products from a store. One approach is to use a hash table to keep track of customer purchases and count the number of unique items purchased. This has an average-case time complexity of O(1) for lookup, but can be expensive to initialize for large datasets. Another approach is to iterate through all customers and check their purchase history, which has a time complexity of O(n). However, this may not be efficient for very large datasets. Ultimately, the choice of approach depends on the size of the dataset and the available resources.

## Time Complexity

O(n) for brute force approach, O(1) for optimized idea.

## Flashcard

Count unique items purchased to find loyal customers.

---

Difficulty: 2