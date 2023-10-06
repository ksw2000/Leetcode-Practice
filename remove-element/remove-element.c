// https://leetcode.com/problems/remove-element/
// Time: O(numsSize)
// Space: O(1)

int removeElement(int* nums, int numsSize, int val){
    int j = 0;
    for(int i = 0; i < numsSize; i++){
        if(nums[i] != val){
            nums[j++] = nums[i];
        }
    }
    return j;
}