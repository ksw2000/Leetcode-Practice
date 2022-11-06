// URL : https://leetcode.com/problems/palindrome-number/
// 8 ms	5.8 MB

#include<stdio.h>
#include<stdlib.h>
typedef enum{
    false, true
}bool;

bool isPalindrome(int x){
    int tmp = x;
    long inverse = 0;
    if(tmp < 0) return false;
    if(tmp < 10) return true;
    while(tmp > 0){
        inverse = inverse * 10 + tmp % 10;
        tmp /= 10;
    }
    return x == inverse;
}

int main(){
    printf("%d\n", isPalindrome(-28));
    printf("%d\n", isPalindrome(8987));
    printf("%d\n", isPalindrome(88));
    printf("%d\n", isPalindrome(878));
    printf("%d\n", isPalindrome(8998));
    printf("%d\n", isPalindrome(1102));
    printf("%d\n", isPalindrome(2147483647));

    return 0;
}
