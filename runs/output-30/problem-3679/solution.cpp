#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minDiscard(vector<int>& weights) {
        unordered_map<int, int> count;
        for (auto w : weights) 
            count[w]++;

        priority_queue<pair<double, int>> pq;
        for (auto &p : count) 
            pq.push({1.0*p.second/weights.size(), p.first});

        int res = 0;
        while (!pq.empty()) {
            auto top = pq.top(); pq.pop();
            if (pq.empty() || top.first > pq.top().first) 
                return -1; 

            res += count[top.second] - ceil(weights.size() * pq.top().first);
        }

        return res;
    }
};