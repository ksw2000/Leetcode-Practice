// https://leetcode.com/problems/two-sum/
// 1 ms	39.9 MB
import java.util.*;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> m = new HashMap<Integer,Integer>();
        int len = nums.length;
        for(int i=0; i<len; i++){
            if(m.containsKey(target - nums[i])){
                return new int[]{m.get(target - nums[i]), i};
            }
            m.put(nums[i], i);
        }
        return new int[]{};
    }
}

public class TwoSum{
    public static void main(String[] args){
        Solution s = new Solution();
        int[] nums = {2, 7, 11, 15};
        int[] ans = s.twoSum(nums, 9);
        for(int v : ans){
            System.out.print(v+" ");
        }
    }
}
