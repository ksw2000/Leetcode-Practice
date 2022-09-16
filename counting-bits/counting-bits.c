// https://leetcode.com/problems/counting-bits/
// 47 ms 10.5 MB
#include <stdio.h>
#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* countBits(int n, int* returnSize) {
    *returnSize = ++n;
    int* ret = malloc(sizeof(int) * n);
    ret[0] = 0;
    int i;
    for (i = 1; i ^ n; i++) {
        ret[i] = ret[i >> 1] + (i & 1);
    }
    return ret;
}

int main() {
    // 0 1 10 11 100 101 110 111 1000 1001 1010
    // 0 1 1 2 1 2 2 3 1 2 2
    int i, size;
    int* arr = countBits(10, &size);

    for (i = 0; i < size; i++) {
        printf("%d%c", arr[i], i != size - 1 ? ' ' : '\n');
    }

    return 0;
}