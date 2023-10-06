// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Time: O(n)

package main

func strStr(haystack string, needle string) int {
out:
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue out
			}
		}
		return i
	}
	return -1
}
