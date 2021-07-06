package main

import (
	"fmt"
	"sort"
	"sync"
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
	finish := make(chan struct{})
	count := 0
	var l sync.Mutex
	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		} else {
			count++
			i2 := i
			go func() {
				for j := i2 + 1; j < length-2; j++ {
					for k := j + 1; k < length-1; k++ {
						search := target - nums[i2] - nums[j] - nums[k]
						if binarySearch(nums, k+1, length-1, search) {
							l.Lock()
							res = append(res, []int{nums[i2], nums[j], nums[k], search})
							l.Unlock()
							// decide next index
							for j+1 < length-2 && nums[j+1] == nums[j] {
								j++
							}
							for k+1 < length-1 && nums[k+1] == nums[k] {
								k++
							}
						}
					}
				}
				finish <- struct{}{}
			}()
		}
	}

	if count > 0 {
		wait := 0
		for range finish {
			wait++
			if wait == count {
				close(finish)
			}
		}
	}

	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 0))
	fmt.Println(fourSum([]int{0, 1, 2}, 0))
	fmt.Println(fourSum([]int{0}, 0))
}
