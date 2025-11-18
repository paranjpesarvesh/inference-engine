Sure, here is a simple solution for this problem in C++17:

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int findSpecialInteger(vector<int>& arr) {
        int n = arr.size();
        for (int i = 0; i < n; ++i) {
            if (count(arr.begin(), arr.end(), arr[i]) > n / 4) {
                return arr[i];
           <｜begin▁of▁sentence｜>}
        }
        return -1;
    }
};

This code defines a class `Solution` with a method `findSpecialInteger` that takes an integer vector as input and returns the element in the array that appears more than 25% of the time. If no such element exists, it returns -1. The function uses the built-in `count` function from the STL to count the number of occurrences of each element in the array.