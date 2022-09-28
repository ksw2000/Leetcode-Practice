// https://leetcode.com/problems/first-missing-positive/
// 60 ms 8.1 MB

package main

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	i := sort.SearchInts(arr, x)
	r := min(i+k, len(arr))
	for ; r > k && abs(arr[r-1]-x) >= abs(arr[r-1-k]-x); r-- {
	}
	return arr[r-k : r]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
