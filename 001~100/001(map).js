// https://leetcode.com/problems/two-sum/
// 44 ms 34.5 MB
var twoSum = function(nums, target) {
    var map=new Object();
    for(let i=0; i<nums.length; i++){
        if(target-nums[i] in map){
            return [i, map[target-nums[i]]];
        }else{
            map[nums[i]]=i;
        }
    }
    return [];
};

var ans1=twoSum([2, 7, 11, 15], 9);
console.log(ans1);
