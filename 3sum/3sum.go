// https://leetcode.com/problems/3sum/
// 33 ms 7.4 MB
package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums)-2 && nums[i] <= 0; {
		for l, r := i+1, len(nums)-1; l < r; {
			cmp := nums[l] + nums[r] + nums[i]
			if cmp > 0 {
				// quick skip
				if nums[r]-nums[l+1] < cmp {
					break
				}
				r--
			} else if cmp < 0 {
				// quick skip
				if nums[l]-nums[r-1] > cmp {
					break
				}
				l++
			} else {
				ans = append(ans, []int{nums[i], nums[r], nums[l]})
				for l++; l <= r && nums[l] == nums[l-1]; l++ {
				}
				for r--; l <= r && nums[r] == nums[r+1]; r-- {
				}
			}
		}
		for i++; i < len(nums) && nums[i] == nums[i-1]; i++ {
		}
	}
	return ans
}
