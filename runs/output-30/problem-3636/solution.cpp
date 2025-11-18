#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<bool> checkIfPrerequisite(int numCourses, vector<vector<int>>& prerequisites, vector<vector<int>>& queries) {
        vector<unordered_set<int>> graph(numCourses);
        for (auto& p : prerequisites) {
            graph[p[0]].insert(p[1]);
        }

        vector<bool> res;
        for (auto& q : queries) {
            int a = q[0], b = q[1];
            if (!graph[a].count(b)) {
                queue<int> q({a});
                unordered_set<int> visited{a};
                while (!q.empty()) {
                    int curr = q.front();
                    q.pop_front();
                    for (auto& next : graph[curr]) {
                        if (next == b) {
                            res.push_back(true);
                            goto end;
                        }
                        if (!visited.count(next)) {
                            visited.insert(next);
                            q.push_back(next);
                        }
                    }
                }
            } else {
                res.push_back(true);
            }
            end:;
        }

        return res;
    }
};