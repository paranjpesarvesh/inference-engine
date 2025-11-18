Sure, here is a simple C++17 solution for the problem "Count Binary Palindromic Numbers". This code counts all binary palindromes from 0 to n.

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countPalindromicNumbers(int n) {
        if (n < 10) return n+1;

        long long res = 10, level = 2; 
        while (true){
            vector<long long> candidates = {level*3 +1 , level*4};

            for(auto &c : candidates) if(c <= n) res++;
            if(candidates[0] > n && candidates[1] > n) break;

            string s = to_string(level);
            reverse(s.begin(), s.end());
            long long nextLevel = level*2 + stoll(s);
            if (nextLevel <= n) {res += 2; level = nextLevel;}
            else break;
        }

        return res;
    }
};

This solution works by generating all binary palindromes and counting them. It starts from 1, then generates the next level of palindrome numbers (by appending a reversed copy of the current number to the end) until it reaches n or exceeds n. For each new generated number, it checks if it is within the range [0,n] and increments the result counter if so.