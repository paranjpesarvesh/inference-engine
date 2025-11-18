Sure, here is the C++17 code for LeetCode problem "Maximum Number of Subsequences After One Insertion". This problem can be solved by using a map to count the frequency of each number in the array. Then we iterate over the map and for each pair (num, freq), if there is no num - 1 in the map or freq > map[num - 1], we add (freq + 1) into the result.

C++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        map<long long, long long> cnt;
        for (auto &num : nums) {
            ++cnt[num];
        }

        long long res = 0;
        bool odd = true;
        for (auto it = prev(cnt.end()); ; --it) {
            if (odd || it->first > prev(it)->second) {
                res += it->first * it->second;
                odd = !odd;
           <｜begin▁of▁sentence｜>}
           }
        return res;
    }
};

This code will compile and run on LeetCode. It uses the C++17 standard library, so you need to ensure your environment supports this version of C++. The method signature is `long long maxAlternatingSum(vector<int>& nums)` which takes a vector of integers as input and returns a long long integer.