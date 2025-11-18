#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<string> findRestaurant(vector<string>& list1, vector<string>& list2) {
        unordered_map<string, int> map;
        for (int i = 0; i < list1.size(); i++) {
            map[list1[i]] = i;
        }

        int minSum = INT_MAX;
        vector<string> res;
        for (int j = 0; j < list2.size() && j <= minSum; j++) {
            if (map.count(list2[j])) {
                if (j + map[list2[j]] < minSum) {
                    minSum = j + map[list2[j]];
                    res.clear();
                    res.push_back(list2[j]);
                } else if (j + map[list2[j]] == minSum) {
                    res.push_back(list2[j]);
                }
            }
        }

        return res;
    }
};