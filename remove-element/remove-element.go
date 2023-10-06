// https://leetcode.com/problems/remove-element/
// Time: O(n) where n is the length of nums
// Space: O(1)

package main

func removeElement(nums []int, val int) int {
	j := 0
	for i := range nums {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
