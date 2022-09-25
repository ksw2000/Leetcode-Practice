// https://leetcode.com/problems/jump-game/
package name

func canJump(nums []int) bool {
	// the length of nums is greater than 1
	reach := 1
	for i := 0; i < reach && reach < len(nums); i++ {
		if nums[i]+i+1 > reach {
			reach = nums[i] + i + 1
		}
	}
	return reach >= len(nums)
}
