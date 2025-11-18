Sure, here's a simple C++17 solution for this problem using `std::sort` and `std::accumulate` from the standard library:

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int findSmallestGreaterThanAverage(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        double avg = accumulate(nums.begin(), nums.end(), 0.0) / nums.size();
        int ans = 1;
        for (auto &n : nums) {
            if (n < avg) continue;
            else if (n == ans) ++ans;
            else return ans;
       <｜begin▁of▁sentence｜>}
       return ans;
    }
};

This code sorts the input vector `nums` in ascending order, calculates its average and then iterates over the sorted array. If a number is less than the average it's skipped, if it equals to the current answer (which starts at 1) it increments the answer by one, otherwise it returns the current answer.