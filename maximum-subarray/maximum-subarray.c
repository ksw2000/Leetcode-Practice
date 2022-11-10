// https://leetcode.com/problems/maximum-subarray/
// 110 ms 12.3 MB

// time: O(n)
int maxSubArray(int* nums, int numsSize) {
    int i;
    int max = nums[0];
    int sum = 0;
    for (i = 0; i < numsSize; i++) {
        sum += nums[i];
        if (sum > max)
            max = sum;
        if (sum < 0)
            sum = 0;
    }

    return max;
}