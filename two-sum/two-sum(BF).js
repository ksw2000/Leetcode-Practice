// https://leetcode.com/problems/two-sum/
// 116 ms 34.7 MB
var twoSum = function(nums, target) {
    var answer=[];
    for(let i=0; i<nums.length-1; i++){
        for(let j=i+1; j<nums.length; j++){
            if(nums[i]+nums[j]===target){
                answer[0]=i;
                answer[1]=j;
                break;
            }
        }
    }
    return answer;
};
