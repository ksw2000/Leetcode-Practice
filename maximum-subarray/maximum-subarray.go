// https://leetcode.com/problems/maximum-subarray/
// 88 ms 4.3 MB

package main

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sumList := []int{0}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sumList = append(sumList, sum)
	}
	max := nums[0]
	for i := 0; i < len(sumList)-1; i++ {
		for j := i + 1; j < len(sumList); j++ {
			if tmp := sumList[j] - sumList[i]; tmp > max {
				max = tmp
			}
		}
	}
	return max
}
