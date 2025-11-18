Here is a C++17 solution for LeetCode problem "Count Bowl Subarrays":

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    long long countSubarrays(vector<int>& nums) {
        int n = nums.size();
        vector<long long> dp(n+1, 0);
        stack<int> stk;

        for (int i = 0; i <= n; ++i) {
            while (!stk.empty() && (i == n || nums[stk.top()] >= nums[i])) {
                int j = stk.top();
                stk.pop();

                long long k = stk.empty() ? -1 : stk.top();
                dp[i] += (nums[j] * (j - k) * (i - j));
            }

            if(i < n){
                stk.push(i);
            }
        }

        long long sum = 0;
        for (int i = 1; i <= n; ++i) {
            sum += dp[i];
        }

        return sum + nums[0];
    }
};

This solution uses a stack to keep track of the indices where the array is non-decreasing. For each index, it calculates how many subarrays can be formed with that index as the right end point. The total number of subarrays for this index is stored in `dp[i]`. Finally, sum up all the dp values to get the final result.