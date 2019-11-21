// URL : https://leetcode.com/problems/remove-duplicates-from-sorted-array
//68 ms	36.8 MB

var removeDuplicates = function(nums) {
    var length=nums.length, j=0;
    for(let i=1; i<length; i++){
        if(nums[i]!=nums[j]){
            nums[++j]=nums[i];
        }
    }
    nums.length=(!length)? 0:j+1;
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
