Sure, here is the C++17 code for LeetCode problem "Find Zombie Sessions". 

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<vector<int>> findZombies(vector<vector<int>>& intervals) {
        sort(intervals.begin(), intervals.end());
        int n = intervals.size();
        if (n == 0) return {};

        vector<vector<int>> res;

        int start = intervals[0][0], end = intervals[0][1];
        for(int i = 1; i < n; ++i){
            if (intervals[i][2] == 1 && intervals[i][0] <= end) { 
                end = max(end, intervals[i][1]);
            } else {  
                res.push_back({start, end});
                start = intervals[i][0];
                end = intervals[i][1];
            }
        }
        res.push_back({start, end}); 
        return res;
    }
};

This code sorts the input intervals by their starting times and then iterates over them to find overlapping intervals (zombie sessions). The result is a list of non-overlapping intervals. Each interval represents a zombie session with its start time at index 0 and end time at index 1.