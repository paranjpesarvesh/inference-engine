#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maximumTotal(vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;
        vector<int> dp(n, 0);
        dp[0] = nums[0];
        for (int i = 1; i < n; ++i) {
            dp[i] = max((i-2>=0?dp[i-2]:0) + nums[i], dp[i-1]);
        }
        return dp.back();
    }
};