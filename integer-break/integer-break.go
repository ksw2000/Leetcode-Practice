// https://leetcode.com/problems/integer-break/
// Space: O(n)
// Time: O(n^2)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	cache := [59]int{0, 0, 1}
	for j := 2; j <= n; j++ {
		m := 1
		for i := 1; i <= j>>1; i++ {
			if k := max(i, cache[i]) * max(j-i, cache[j-i]); k > m {
				m = k
			}
		}
		cache[j] = m
	}
	return cache[n]
}