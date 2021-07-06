package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, start, end, target int) bool {
	if target > nums[end] || target < nums[start] {
		return false
	}
	for start != end {
		next := (start + end) >> 1
		if target > nums[next] {
			start = next + 1
		} else if target < nums[next] {
			end = next
		} else {
			return true
		}
	}
	return target == nums[start]
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			for k := j + 1; k < length-1; k++ {
				search := target - nums[i] - nums[j] - nums[k]
				if binarySearch(nums, k+1, length-1, search) {
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
	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 0))
}
