// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// 36 ms 2 MB

package main

func concatenatedBinary(n int) int {
	var sum int
	var j int8
	n++
	for i := 1; i < n; i++ {
		if (i & (i - 1)) == 0 {
			j++
		}
		sum = (sum<<j | i) % 1000000007
	}
	return sum
}
