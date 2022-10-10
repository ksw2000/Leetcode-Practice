// https://leetcode.com/problems/break-a-palindrome/
// 0 ms 2 MB

package main

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}
	mutable := []byte(palindrome)
	// change the first character which is not 'a'
	// e.g.
	// abccba -> aaccba
	// azaza -> aaaza
	for i := 0; i < len(palindrome)>>1; i++ {
		if palindrome[i] != 'a' {
			mutable[i] = 'a'
			return string(mutable)
		}
	}
	// if palindrome is aaa...aaa, or aabaa
	// change the last character
	mutable[len(mutable)-1] = 'b'
	return string(mutable)
}
