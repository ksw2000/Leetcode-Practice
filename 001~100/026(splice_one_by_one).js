// URL : https://leetcode.com/problems/remove-duplicates-from-sorted-array
//92 ms	37.3 MB

var removeDuplicates = function(nums) {
    var length=nums.length, last, i=0;
    last=nums[i++];
    for(; i<length; i++){
        if(nums[i]==last){
            nums.splice(i,1);
            i--;
            length--;
        }
        last=nums[i];
    }
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
