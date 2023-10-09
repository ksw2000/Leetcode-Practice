// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Time: O(lg n) where n is the length of nums
// Space: O(1)

package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l := 0
	r := len(nums) - 1
	var m int
	for l <= r {
		m = (l + r) >> 1
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = l + 1
		} else {
			break
		}
	}

	if nums[m] != target {
		return []int{-1, -1}
	}

	// find left
	l = m
	for l >= 0 && nums[l] == target {
		step := 1
		for l-step<<1 >= 0 && nums[l-step<<1] == target {
			step = step << 1
		}
		l = l - step
	}

	// find right
	r = m
	for r < len(nums) && nums[r] == target {
		step := 1
		for r+step<<1 < len(nums) && nums[r+step<<1] == target {
			step = step << 1
		}
		r = r + step
	}

	return []int{l + 1, r - 1}
}
