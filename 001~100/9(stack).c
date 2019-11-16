// URL : https://leetcode.com/problems/palindrome-number/
#include<stdio.h>
#include<stdlib.h>
typedef enum{
    false, true
}bool;

bool isPalindrome(int x){
    if(x<0) return false;
    int tmp=x,stack[14],index=-1;

    while(tmp>0){
        stack[++index]=tmp%10;
        tmp/=10;
    }

    while(x>0){
        if(x%10 != stack[index--]){
            return false;
        }
        x/=10;
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
