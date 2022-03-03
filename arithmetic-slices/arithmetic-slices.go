// https://leetcode.com/problems/arithmetic-slices/
// 0 ms	2.1 MB

package main

// [1 2 3 4 5 7]
//  ~~^~~			j = 1 	count = 1
// [1 2 3 4 5 7]
//    ~~^~~			j = 2	count = 3
// [1 2 3 4 5 7]
//      ~~^~~       j = 3	count = 6
// [1 2 3 4 5 7]
//        ~~^~~     j = 0	count = 6

func numberOfArithmeticSlices(nums []int) int {
	count := 0
	for i, j := 1, 0; i+1 < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i+1]-nums[i] {
			j++
			count += j
		} else {
			j = 0
		}
	}
	return count
}
