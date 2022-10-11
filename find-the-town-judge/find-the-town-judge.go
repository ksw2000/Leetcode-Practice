// https://leetcode.com/problems/find-the-town-judge/
// 96 ms 7.3 MB
package main

func findJudge(n int, trust [][]int) int {
	counter := make([]int, n)
	for i := range trust {
		counter[trust[i][1]-1]++
		counter[trust[i][0]-1]--
	}
	for i := range counter {
		if counter[i] == n-1 {
			return i + 1
		}
	}
	return -1
}
