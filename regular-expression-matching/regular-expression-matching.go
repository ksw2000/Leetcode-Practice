// https://leetcode.com/problems/regular-expression-matching/
// 8 ms 2 MB
package main

func isMatchRune(s []rune, p []rune) bool {
	for len(p) > 1 {
		// syntax error
		// if p[0] == '*' {
		// 	return false
		// }

		if p[0] == '.' {
			if p[1] == '*' {
				// if partern is .*
				ret := false
				for j := 0; j <= len(s); j++ {
					ret = ret || isMatchRune(s[j:], p[2:])
				}
				return ret
			}
			// partern is . but string is empty
			if len(s) == 0 {
				return false
			}
		} else {
			if p[1] == '*' {
				// if partern is a* b* c* ...
				ret := false
				j := 0
				for j < len(s) && s[j] == p[0] {
					ret = ret || isMatchRune(s[j:], p[2:])
					j++
				}
				return ret || isMatchRune(s[j:], p[2:])
			}
			if len(s) == 0 || p[0] != s[0] {
				return false
			}
		}
		s = s[1:]
		p = p[1:]
	}

	if len(p) > 0 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}
	return len(s) == 0
}

func isMatch(s string, p string) bool {
	return isMatchRune([]rune(s), []rune(p))
}
