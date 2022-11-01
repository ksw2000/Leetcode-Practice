// https://leetcode.com/problems/pascals-triangle-ii/
// 0 ms 5.9 MB
// Time complexity = O(rowIndex/2)

#include <stdio.h>
#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* getRow(int rowIndex, int* returnSize) {
    // [ 1 ] [  ] ... [ l+rowIndex ]
    //   l        ...        r
    int* l = malloc(sizeof(int) * (rowIndex + 1));
    int* r = l + rowIndex;
    l[0] = r[0] = 1;
    *returnSize = (rowIndex + 1);

    int i;
    for (i = 1; i <= (rowIndex >> 1); i++) {
        l[i] = r[-i] = (unsigned long long)(l[i - 1]) * (rowIndex - i + 1) / i;
    }
    return l;
}