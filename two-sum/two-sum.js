// https://leetcode.com/problems/two-sum/
// 44 ms 34.5 MB
var twoSum = function(nums, target) {
    var map = {};
    var len = nums.length;
    for(let i=0; i<len; i++){
        if(map[target-nums[i]] !== undefined){
            return [i, map[target-nums[i]]];
        }
        map[nums[i]]=i;
    }
    return [];
};

var ans1=twoSum([2, 7, 11, 15], 9);
console.log(ans1);
