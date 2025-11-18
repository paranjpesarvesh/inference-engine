#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int largestPermutation(int k, vector<int>& arr) {
        int n = arr.size();
        unordered_map<int, int> m; 

        for (int i = 0; i < n; ++i) 
            m[arr[i]] = i;

        for (int i = 0; i < n && k > 0; ++i) {
            if (arr[i] != n - i) { 
                int temp = arr[i]; 
                swap(arr[i], arr[m[n-i]] ); 

                m[temp] = m[n - i]; 
                m[n-i] = i; 

                --k; 
            }
       <｜begin▁of▁sentence｜>
	   }
       return n-k;
    }
};