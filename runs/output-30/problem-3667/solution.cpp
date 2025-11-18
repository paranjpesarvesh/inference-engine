#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        int n = nums.size();
        vector<int> result(n);
        int i = 0, j = n - 1, k = n - 1;

        while (i <= j) {
            if (abs(nums[i]) > abs(nums[j])) {
                result[k--] = nums[i] * nums[i];
                ++i;
            } else {
                result[k--] = nums[j] * nums[j];
                --j;
            }
        }

        return result;
    }
};