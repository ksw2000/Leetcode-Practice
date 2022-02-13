// https://leetcode.com/problems/simplify-path/
// 0 ms 6.3 MB
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char** split(char* path, int* retLen) {
    int i = 0;
    int j = 0;
    int ret_len = 0;
    int ret_cap = 8;
    char** ret = malloc(sizeof(char*) * ret_cap);
    int breakingFlag = 1;
    while (breakingFlag) {
        if (path[i] == '/' || path[i] == '\0') {
            if (path[i] == '\0') {
                breakingFlag = 0;
            }
            if (i - j > 0) {
                if (ret_len + 1 > ret_cap) {
                    // enlarge
                    ret_cap = ret_cap << 1;
                    char** newRet = malloc(sizeof(char*) * ret_cap);
                    int k = 0;
                    for (k = 0; k < ret_len; k++) {
                        newRet[k] = ret[k];
                    }
                    free(ret);
                    ret = newRet;
                }
                ret[ret_len++] = path + j;
                path[i] = '\0';
            }
            j = i + 1;
        }
        i++;
    }
    *retLen = ret_len;
    return ret;
}

char* simplifyPath(char* path) {
    int i = 0;  // for loop
    int j = 0;  // for adjusting pathList
    int pathListLen;
    char** pathList = split(path, &pathListLen);
    for (i = 0; i < pathListLen; i++) {
        if (!strcmp(pathList[i], "..")) {
            if (j > 0)
                j--;
        } else if (strcmp(pathList[i], "") && strcmp(pathList[i], ".")) {
            pathList[j] = pathList[i];
            j++;
        }
    }

    // concatenate string in the pathList
    char* c;
    int k = 0;
    if (j > 0) {
        for (i = 0; i < j; i++) {
            path[k++] = '/';
            c = pathList[i];
            while (*c != '\0') {
                path[k++] = *c;
                c++;
            }
        }
        path[k++] = '\0';
        return path;
    }
    return "/";
}