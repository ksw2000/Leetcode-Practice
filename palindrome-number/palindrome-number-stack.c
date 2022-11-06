// URL : https://leetcode.com/problems/palindrome-number/
// 12 ms 5.9 MB
#include<stdio.h>
#include<stdlib.h>
typedef enum{
    false, true
}bool;

bool isPalindrome(int x){
    if(x < 0) return false;
    if(x < 10) return true;
    int index = -1;
    int stack[10];

    while(x > 0){
        stack[++index] = x % 10;
        x /= 10;
    }

    int half = (index + 1) >> 1;
    for(; half >= 0; half--){
        if(stack[index-half] != stack[half]) return false;
    }
    return true;
}

int main(){
    printf("%d\n",isPalindrome(-28));
    printf("%d\n",isPalindrome(0));
    printf("%d\n",isPalindrome(1));
    printf("%d\n",isPalindrome(8987));
    printf("%d\n",isPalindrome(88));
    printf("%d\n",isPalindrome(878));
    printf("%d\n",isPalindrome(8998));
    printf("%d\n",isPalindrome(1102));
    printf("%d\n",isPalindrome(2147483647));

    system("pause");
    return 0;
}
