// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// 11 ms 7.8 MB

#include<stdio.h>
int removeDuplicates(int* nums, int numsSize) {
    int i;
    int j = 1;
    for (i = 1; i < numsSize; i++) {
        if (nums[i] != nums[i - 1]) {
            nums[j] = nums[i];
            j++;
        }
    }
    return j;
}