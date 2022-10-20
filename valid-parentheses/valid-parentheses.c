// https://leetcode.com/problems/valid-parentheses/
// 0 ms 5.8 MB

#include<stdio.h>
#include<stdlib.h>

typedef enum{
    false,
    true
}bool;

char convert(char s) {
    if (s == ')')
        return '(';
    if (s == ']')
        return '[';
    if (s == '}')
        return '{';
    return 0;
}

bool isValid(char* s) {
    int i = 0;
    int j = 0;
    int token;
    for (i = 0; s[i] != '\0'; i++) {
        token = convert(s[i]);
        if (token == 0) {
            s[j] = s[i];
            j++;
        } else if (j > 0 && s[j - 1] == token) {
            j--;
        } else {
            return false;
        }
    }

    return j == 0;
}