#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    string filter(string s, vector<int>& nums) {
        unordered_map<char, int> freq;
        for (auto c : s) {
            freq[c]++;
        }

        string result = "";
        for (int i = 0; i < nums.size(); ++i) {
            if (freq[nums[i]] > 0) {
                result += to_string(nums[i]);
            }
        }

        return result;
    }
};