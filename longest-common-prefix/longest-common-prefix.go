// https://leetcode.com/problems/longest-common-prefix/
// 0 ms	2.3 MB

package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	limit := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < limit {
			limit = len(strs[i])
		}
	}

	max := 0
l1:
	for i := 0; i < limit; i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				break l1
			}
		}
		max++
	}
	return strs[0][:max]
}
