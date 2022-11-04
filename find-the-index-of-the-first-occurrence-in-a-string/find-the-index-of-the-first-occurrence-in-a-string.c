// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// 0 ms 7.4 MB
#include <stdio.h>
int strStr(char* haystack, char* needle) {
    int needle_len = strlen(needle);
    int max_i = strlen(haystack) - needle_len;

    if (needle_len == 0) {
        return 0;
    }

    int i, j, tmp_i;
    for (i = 0; i <= max_i; i++) {
        if (haystack[i] == needle[0]) {
            tmp_i = i;
            for (j = 0; j < needle_len; j++) {
                if (haystack[tmp_i++] != needle[j]) {
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