// https://leetcode.com/problems/jump-game-ii/
// 16 ms 6.2 MB

package main

func jump(nums []int) int {
	reach := 0
	jump := 0
	update := 0
	for i := 0; update < len(nums)-1; i++ {
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
		if i == update {
			jump++
			update = reach
		}
	}
	return jump
}
