// https://leetcode.com/problems/maximum-subarray/
// 106 ms 8.8 MB

package main

func maxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i := range nums {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
