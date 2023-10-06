// https://leetcode.com/problems/permutation-sequence/
// Time: O(n^2)

package main

func getPermutation(n int, k int) string {
	ans := make([]byte, n)
	f := 1
	for i := 1; i <= n; i++ {
		ans[i-1] = byte(i) + '0'
		f *= i
	}

	k--
	for i := 0; i < n; i++ {
		f /= n - i
		j := k/f + i
		tmp := ans[j]
		for k := j; k > i; k-- {
			ans[k] = ans[k-1]
		}
		ans[i] = tmp

		k = k % f
	}
	return string(ans)
}
