package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			k := j + 1
			l := length - 1
			search := target - nums[i] - nums[j]
			for l > k {
				try := nums[k] + nums[l]
				if try > search {
					l--
					for nums[k]+nums[(k+l)>>1] > search {
						if l == (k+l)>>1 {
							break
						}
						l = (k + l) >> 1
					}
				} else if try < search {
					k++
					for nums[(k+l)>>1]+nums[l] < search {
						if k == (k+l)>>1 {
							break
						}
						k = (k + l) >> 1
					}
				} else {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					for l--; l > j+1 && nums[l] == nums[l+1]; l-- {
					}
					for k++; k < length-1 && nums[k] == nums[k-1]; k++ {
					}
				}
			}
			for j+1 < length-2 && nums[j+1] == nums[j] {
				j++
			}
		}
		for i+1 < length-3 && nums[i+1] == nums[i] {
			i++
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
