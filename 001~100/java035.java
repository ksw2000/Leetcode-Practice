// URL : https://leetcode.com/problems/search-insert-position/
//0 ms	39 MB
class Solution {
    public int searchInsert(int[] nums, int target) {
        int i;
        for(i=0; i<nums.length; i++){
            if(nums[i] >= target){
                return i;
            }
        }
        return nums.length;
    }
}

public class java035{
    public static void main(String args[]){
        Solution sol=new Solution();
        int[] array={1,3,5,6};
        System.out.println(sol.searchInsert(array,5));
        System.out.println(sol.searchInsert(array,7));
        System.out.println(sol.searchInsert(array,0));
        int[] array2={};
        System.out.println(sol.searchInsert(array2,10));
    }
}
