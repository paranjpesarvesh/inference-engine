#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<vector<int>> decompose(int n, int k) {
        if (n < 0 || k < 1) return {};

        vector<vector<int>> res(k, vector<int>(1, 1));

        for (int i = 0; i < n - k && res[i % k].size() <= n / k + 1; ++i) {

            ++res[i % k][0];
        }

        return res;
    }
};