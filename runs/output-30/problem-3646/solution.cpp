Sure, here's a simple C++17 solution for this problem. This code increments the given number and checks if it's a palindrome. If not, it continues to increment until it finds a palindrome. 

C++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    string nextPalindrome(string num) {
        int n = num.size();
        int mid = (n-1)/2;

        bool greater_half = false;
        for(int i=mid;i>=0;i--){
            if(num[i]<num[n-1-i]) {
                greater_half = true;
                break;
            } else if (num[i]>num[n-1-i]) {
                reverse(num.begin()+i, num.end());
                return num;
           <｜begin▁of▁sentence｜>}
        }

        if(!greater_half) {
            reverse(num.begin(), num.end());
            return num;
        } else {
            int i = mid;
            while(i>=0 && num[i] >= num[n-1-i]) {
                num[i] = '0';
                if(i>0) num[n-1-i--] = '0';
            }

            if(i>=0){
                num[i]++;
                int j=n-1-i;
                while(j<n && i>=0){
                    num[j] = num[i];
                    j++, i--;
                }
            } else {
                reverse(num.begin(), num.end());
                return "1" + string(n-1, '0') + "1";
            }
        }

        while(!num.empty() && num[0]=='0'){
            num = num.substr(
}
}
}