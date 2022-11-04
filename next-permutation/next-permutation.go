// https://leetcode.com/problems/next-permutation/
// 0 ms 2.5 MB
package main

import "sort"

func nextPermutation(nums []int) {
	// find increasing-order subarray from right to left
	// 1 3 4 2
	//    ^
	// find the minimum value, which is larger than the
	// rightest digit of the left part, in the left part
	//
	// 1 4 3 2
	//    ^
	// sort the right part in increasing order from left to right
	// 1 4 2 3

	// input: 1 2 4 3
	//           ^
	// swap:  1 3 4 2
	//           ^
	// sort:  1 3 2 4

	// find breaking point

	b := len(nums) - 1
	for ; b >= 1 && nums[b] <= nums[b-1]; b-- {
	}
	if b != 0 {
		// find the minimum value larger than nums[b-1]
		c := b - 1 + find(nums[b:], nums[b-1])
		// fmt.Println("breaking", b, "c:", c)
		// swap
		nums[b-1], nums[c] = nums[c], nums[b-1]
	}
	sort.Ints(nums[b:])
}

// find the minimum number larget than target
func find(nums []int, target int) int {
	current := 0
	skip := len(nums) >> 1
	for current < len(nums) && nums[current] > target {
		if nums[current+skip] > target && skip > 0 {
			current += skip
		} else {
			current++
		}
		skip = skip >> 1
	}
	return current
}
