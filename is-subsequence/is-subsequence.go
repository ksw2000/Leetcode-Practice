// https://leetcode.com/problems/is-subsequence/
// 0 ms 2 MB

package main

import "fmt"

func isSubsequence(s string, t string) bool {
	ss, tt := []byte(s), []byte(t)
	var i, j int
	for i < len(ss) && j < len(tt) {
		if ss[i] == tt[j] {
			i++
		}
		j++
	}
	return i == len(ss)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
