package main

import (
	"math"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	minD := math.MaxInt32
	closest := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		d := target - nums[i] - nums[j] - nums[k]
		for k > j+1 {
			if d < 0 {
				k--
			} else if d > 0 {
				j++
			} else {
				return target
			}
			d = target - nums[i] - nums[j] - nums[k]
		}
		if abs(d) < minD {
			minD = abs(d)
			closest = target - d
		}
	}
	return closest
}
