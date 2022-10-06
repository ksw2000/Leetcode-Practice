// https://leetcode.com/problems/first-missing-positive/
// 36 ms 8.3 MB

package main

func firstMissingPositive(nums []int) int {
	j := len(nums)
	i := 0
	for i < j {
		if nums[i] != i+1 {
			if nums[i] > 0 && nums[i] <= j && nums[nums[i]-1] != nums[i] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			} else {
				j--
				nums[i] = nums[j]
			}
		} else {
			i++
		}
	}
	return i + 1
}
