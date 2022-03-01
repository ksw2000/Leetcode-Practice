// https://leetcode.com/problems/counting-bits/
// 2 ms	4.5 MB
package main

// This is a very simple solution
// Time complexity is O(n log n)

func countBits(n int) []int {
	ret := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		for j := i; j > 0; j = j >> 1 {
			ret[i] += j & 1
		}
	}
	return ret
}
