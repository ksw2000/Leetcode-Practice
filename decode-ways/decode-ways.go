// https://leetcode.com/problems/decode-ways/
// 0 ms	1.9 MB
package main

func numDecodings(s string) int {
	prepre := 1
	pre := 0
	if len(s) > 0 && s[0] != '0' {
		pre = 1
	}
	for i := 1; i < len(s); i++ {
		sum := 0
		if s[i] != '0' {
			sum += pre
		}
		if n := (s[i-1]-'0')*10 + s[i] - '0'; n >= 10 && n <= 26 {
			sum += prepre
		}
		prepre = pre
		pre = sum
	}
	return pre
}
