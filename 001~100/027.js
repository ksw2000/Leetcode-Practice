// URL : https://leetcode.com/problems/remove-element/
//52 ms	33.7 MB

var removeElement = function(nums, val) {
    var length=nums.length;
    for(let i=0; i<length; i++){
        if(nums[i]==val){
            nums.splice(i,1);
            i--;
            length--;
        }
    }
    return length;
};

var nums=[10,9,8,3,3,3,4,5,9];
console.log(removeElement(nums,3));
console.log(nums);
