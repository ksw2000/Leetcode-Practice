// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// 3 ms 5.9 MB
#include <stdio.h>

#define invertLen 256
int lengthOfLongestSubstring(char* s) {
    int invert[invertLen];  // char -> index
    int i = 0;
    for (i = 0; i < invertLen; i++) {
        invert[i] = 0;
    }
    int start = 0;
    int max = 0;
    for (i = 0; s[i] != '\0'; i++) {
        if (invert[s[i]] > start) {
            if (i - start > max) {
                max = i - start;
            }
            start = invert[s[i]];
        }
        invert[s[i]] = i + 1;
    }
    if (i - start > max) {
        max = i - start;
    }
    return max;
}