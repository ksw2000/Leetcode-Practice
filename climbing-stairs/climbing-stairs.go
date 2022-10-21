// https://leetcode.com/problems/climbing-stairs/
// 0 ms 1.9 MB
package main

import "math"

func climbStairs(n int) int {
	// f(1) = 1
	// f(2) = 2
	// f(n) = f(n-1) + f(n-2)
	n++
	a := math.Sqrt(5)
	return int(math.Round((math.Pow((1+a)/2, float64(n)) - math.Pow((1-a)/2, float64(n))) / a))
}
