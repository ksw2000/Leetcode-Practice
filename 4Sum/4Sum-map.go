package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	inverse := map[int]struct{}{}
	for _, num := range nums {
		inverse[num] = struct{}{}
	}
	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			for k := j + 1; k < length-1; k++ {
				search := target - nums[i] - nums[j] - nums[k]
				if search > nums[k] || (search == nums[k] && search == nums[k+1]) {
					if _, ok := inverse[search]; ok {
						res = append(res, []int{nums[i], nums[j], nums[k], search})
						// decide next index
						for i+1 < length-3 && nums[i+1] == nums[i] {
							i++
						}
						for j+1 < length-2 && nums[j+1] == nums[j] {
							j++
						}
						for k+1 < length-1 && nums[k+1] == nums[k] {
							k++
						}
					}
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 0))
}
