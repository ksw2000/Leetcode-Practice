// https://leetcode.com/problems/counting-bits/
// 0 ms	4.6 MB
package main

import "fmt"

// The fastest solution
// Time complexity O(n)

func countBits(n int) []int {
	ret := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		ret[i] = ret[i>>1] + i&1
	}
	return ret
}

func main() {
	// 0 1 10 11 100 101 110 111 1000 1001 1010
	// 0 1 1 2 1 2 2 3 1 2 2
	fmt.Println(countBits(10))
}
