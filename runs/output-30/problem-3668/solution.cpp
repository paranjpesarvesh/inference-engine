#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> arrayRankTransform(vector<int>& arr) {
        if (arr.empty()) return {};

        auto sorted_arr = arr;
        sort(sorted_arr.begin(), sorted_arr.end());

        unordered_map<int, int> rank;
        for (auto num : sorted_arr) {
            if (!rank.count(num)) {
                rank[num] = rank.size() + 1;
            }
        }

        vector<int> res;
        for (auto num : arr) {
            res.push_back(rank[num]);
        }

        return res;
    }
};