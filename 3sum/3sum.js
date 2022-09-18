// https://leetcode.com/problems/3sum/
// 174 ms 52.4 MB
/**
 * @param {number[]} nums
 * @return {number[][]}
 */

var threeSum = function (nums) {
    nums.sort(function (a, b) {
        return a - b;
    });
    const res = [];
    const len = nums.length;
    let l, r;
    for (let i = 0; i < len - 2;) {
        l = i + 1;
        r = len - 1;
        while (l < r) {
            let sum = nums[i] + nums[l] + nums[r];
            if (sum > 0) {
                for (r--; r >= l && nums[r + 1] === nums[r]; r--);
            } else if (sum < 0) {
                for (l++; r >= l && nums[l - 1] === nums[l]; l++);
            } else {
                res.push([nums[i], nums[l], nums[r]])
                for (l++; r >= l && nums[l - 1] === nums[l]; l++);
                for (r--; r >= l && nums[r + 1] === nums[r]; r--);
            }
        }
        for (i++; i < len && nums[i] === nums[i - 1]; i++);
    }
    return res;
};