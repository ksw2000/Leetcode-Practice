// https://leetcode.com/problems/pascals-triangle/
// 3 ms 6.6 MB
// Use malloc() only 3 times.

#include <stdio.h>
#include <stdlib.h>

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume
 * caller calls free().
 */

int** generate(int numRows, int* returnSize, int** returnColumnSizes) {
    int i, j;
    *returnSize = numRows;
    *returnColumnSizes = malloc(sizeof(int) * numRows);
    int* list = malloc(sizeof(int) * (1 + numRows) * numRows >> 1);
    int** ret = malloc(sizeof(int*) * numRows);

    for (i = 0; i < numRows; i++) {
        int shift = (i * (i + 1)) >> 1;
        ret[i] = list + shift;
        list[shift] = 1;
        list[shift + i] = 1;

        for (j = shift + 1; j < shift + i; j++) {
            list[j] = list[j - i] + list[j - i - 1];
        }

        (*returnColumnSizes)[i] = i + 1;
    }

    return ret;
}