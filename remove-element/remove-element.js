// https://leetcode.com/problems/remove-element/
// Time: O(n) where n is the length of nums
// Space: O(1)

/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
    let j = 0;
    const n = nums.length;
    for(let i = 0; i < n; i++){
        if(nums[i] != val){
            nums[j++] = nums[i];
        }
    }
    return j;
};