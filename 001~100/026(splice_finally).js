// URL : https://leetcode.com/problems/remove-element/
//72 ms	36.6 MB

var removeDuplicates = function(nums) {
    var length=nums.length, last, i=1, j=1;
    last=nums[0];
    for(; i<length; i++){
        if(nums[i]!=last){
            nums[j++]=nums[i];
        }
        last=nums[i];
    }
    nums.splice(j,length-j);
}

var nums=[1,2,2,3,4,5,5,6,6];
removeDuplicates(nums)
console.log(nums);
nums=[];
removeDuplicates(nums)
console.log(nums);
nums=[0];
removeDuplicates(nums)
console.log(nums);
nums=[1,1,2];
removeDuplicates(nums)
console.log(nums);
