// https://leetcode.com/problems/wildcard-matching/
// 4 ms 2.8 MB
package main

// s = ^abcbd$
// p = ^a*c*d$*
//     ^^ ^ ^^
//     p0 p1 p2
// ---------------------
// p0 = ^a   s[0] = ^a
// p1 = c    s[3] = c
// p2 = d$   s[5] = d$
// ---------------------
// return true

// s = ^abcababab$
// p = ^abc*$*
//     ^^^^ ^
//      p0  p1
// ---------------------
// p0 = ^abc  s[0] = ^abc
// p1 = $     s[10] = $
// ---------------------
// return true

// s = ^abc$
// p = ^a$*
//     ^^^
//      p0
// ---------------------
// p0 = ^a$  there is no ^a$ in s
// ---------------------
// return false

func isMatch(s string, p string) bool {
	// ^, $ for matching the first character and the last character
	// the last * in the `p` is used to split `p`
	ss := []byte("^" + s + "$")
	pp := []byte("^" + p + "$*")

	j := 0
	for i := 0; i < len(pp); i++ {
		if pp[i] != '*' {
			continue
		}
		// for duplicate *
		if i == j {
			j = i + 1
			continue
		}
		if k := firstMatch(ss, pp[j:i]); k != -1 {
			ss = ss[k+i-j:]
		} else {
			return false
		}
		j = i + 1
	}
	return true
}

// (1)
// s = abcabc
//     ^^^
// p = abc
// return 0

// (2)
// s = ababc
//       ^^^
// p = abc
// return  2

// (3)
// s = aabbc
//       ^^^
// p = b?c
// return 2

// (4)
// s = acbde
// p = ab
// return -1

func firstMatch(s, p []byte) int {
outer:
	for i := 0; i < len(s)-len(p)+1; i++ {
		for j := 0; j < len(p); j++ {
			if s[i+j] != p[j] && p[j] != '?' {
				continue outer
			}
		}
		return i
	}
	return -1
}
