// https://leetcode.com/problems/valid-parentheses/
// 0 ms 1.8 MB
package main

func convert(s byte) byte {
	if s == ')' {
		return '('
	}
	if s == ']' {
		return '['
	}
	if s == '}' {
		return '{'
	}
	return 0
}

func isValid(s string) bool {
	ss := []byte(s)
	j := 0
	for i := range ss {
		token := convert(ss[i])
		if token == 0 {
			ss[j] = ss[i]
			j++
		} else if j > 0 && ss[j-1] == token {
			j--
		} else {
			return false
		}
	}
	return j == 0
}
