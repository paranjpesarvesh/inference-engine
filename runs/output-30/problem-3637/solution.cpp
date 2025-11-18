#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int sumOfArray(vector<int>& nums) {
        return accumulate(nums.begin(), nums.end(), 0);
    }
};