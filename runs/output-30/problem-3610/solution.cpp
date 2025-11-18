Sure, I can help with that. Here's a simple C++17 solution for your problem using Sieve of Eratosthenes to generate prime numbers and then applying dynamic programming to find the minimum number of primes needed to sum up to the target. 

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minPrimeSet(vector<int>& nums, int target) {
        vector<bool> prime(target + 1, true);
        prime[0] = prime[1] = false;
        for (int p = 2; p * p <= target; ++p) {
            if (prime[p]) {
                for (int i = p * 2; i <= target; i += p) {
                    prime[i] = false;
                }
            }
        }

        vector<int> dp(target + 1, INT_MAX);
        dp[0] = 0;
        for (auto num : nums) {
            if (prime[num]) {
                for (int i = target; i >= num; --i) {
                    if (dp[i - num] != INT_MAX) {
                        dp[i] = min(dp[i], dp[i - num] + 1);
                    }
                }
           <｜begin▁of▁sentence｜>}
        }

        return dp[target] == INT_MAX ? -1 : dp[target];
    }
};

This code first generates a boolean array `prime` where `prime[i]` is true if and only if `i` is prime. Then it uses dynamic programming to find the minimum number of primes needed to sum up to each target from 0 to `target`. If no combination can sum up to the target, it returns -1.