// https://leetcode.com/problems/find-the-town-judge/
// 178 ms 16.3 MB
#include <stdio.h>
#include <stdlib.h>

int findJudge(int n, int** trust, int trustSize, int* trustColSize) {
    int i;
    int* counter = calloc(n, sizeof(int));
    for (i = 0; i < trustSize; i++) {
        counter[trust[i][1] - 1]++;
        counter[trust[i][0] - 1]--;
    }
    for (i = 0; i < n; i++) {
        if (counter[i] == n - 1) {
            // free(counter);
            return i + 1;
        }
    }
    // free(counter);
    return -1;
}