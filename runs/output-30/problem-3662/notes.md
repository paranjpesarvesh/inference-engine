# Filter Characters by Frequency

## Description

Given a string of characters and a frequency array, filter the characters based on their frequencies.

## Approaches

• Brute force: Iterate through each character in the string and check if its frequency is present in the frequency array. If it is, add it to the filtered string.
• Optimized idea: Use a dictionary to store the characters as keys and their frequencies as values. Then iterate through the frequency array and add the characters with non-zero frequencies to the dictionary. Finally, convert the dictionary back into a list of characters.
• Key tradeoff: The optimized idea has a time complexity of O(n+k), where n is the length of the string and k is the length of the frequency array. This is faster than the brute force approach, which has a time complexity of O(n^2). However, it requires more memory to store the dictionary.

## Explanation

First, we define two functions: `filter_by_frequency` and `count_chars`. The `filter_by_frequency` function takes a string and a frequency array as input and returns a new string containing only the characters that appear in the frequency array. The `count_chars` function takes a string as input and returns a dictionary where each key is a character and its corresponding value is the number of times it appears in the string.
Next, we use the `filter_by_frequency` function to filter out the characters that do not appear in the frequency array. We then use the `count_chars` function to count the frequencies of each character in the filtered string. Finally, we use a loop to iterate through the frequency array and add the characters with non-zero frequencies to the dictionary.
To convert the dictionary back into a list of characters, we use another loop that iterates through the dictionary and adds each character to a new list. We then return this list as the final output.

## Time Complexity

The time complexity of this solution is O(n+k), where n is the length of the string and k is the length of the frequency array.

## Flashcard

Filter characters by frequency: Count character frequencies, filter out non-frequent characters.

---

Difficulty: 1 • Paid: yes