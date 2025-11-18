#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minOperationsToBalance(vector<int>& nums) {
        int n = nums.size();
        vector<int> prefixSum(n + 1, 0);
        for (int i = 0; i < n; ++i) {
            prefixSum[i+1] = prefixSum[i] + nums[i];
        }

        int totalSum = prefixSum.back();
        if(totalSum % n != 0){
            return -1; 
        }

        int target = totalSum / n;
        int ans = 0, left_sum = 0;
        for (int i = 0; i < n; ++i) {
            int excess = left_sum + nums[i] - target;
            if(excess > 0){

                ans += excess;

                left_sum += excess; 
            }
       <｜begin▁of▁sentence｜>	left_sum += nums[i];
}
}
}