// URL : https://leetcode.com/problems/implement-strstr/
//0 ms	7.4 MB
#include<stdio.h>
#include<stdlib.h>
#include<string.h>

int strStr(char * haystack, char * needle){
    int needle_len=strlen(needle);
    int max_i=strlen(haystack)-needle_len;

    if(needle_len==0){
        return 0;
    }

    int i,j,tmp_i;
    for(i=0; i<=max_i; i++){
        if(haystack[i]==needle[0]){
            tmp_i=i;
            for(j=0; j<needle_len; j++){
                if(haystack[tmp_i++]!=needle[j]){
                    goto CONTINUE;
                }
            }
            return i;
        }
        CONTINUE:
        continue;
    }
    return -1;
}

int main(){
    printf("%d\n",strStr("hello","ll"));            // 2
    printf("%d\n",strStr("llo","ll"));              // 0
    printf("%d\n",strStr("lo","ll"));               // -1
    printf("%d\n",strStr("l","l"));                 // 0
    printf("%d\n",strStr("",""));                   // 0
    printf("%d\n",strStr("hello",""));              // 0
    printf("%d\n",strStr("","hello"));              // -1
    printf("%d\n",strStr("mississippi","issip"));   // 4
    system("pause");
    return 0;
}
