<?php
// https://leetcode.com/problems/two-sum/
// 8 ms	15.9 MB

class Solution {
    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $map = Array();
        $len = count($nums);
        for($i=0; $i<$len; $i++){
            if(isset($map[$target - $nums[$i]])){
                return [$map[$target - $nums[$i]], $i];
            }
            $map[$nums[$i]] = $i;
        }
        return [];
    }
}

$nums = Array(2, 7, 11, 15);
$s = new Solution();

var_dump($s->twoSum($nums, 9));

?>
