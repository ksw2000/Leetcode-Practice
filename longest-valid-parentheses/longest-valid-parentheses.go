// https://leetcode.com/problems/longest-valid-parentheses/
// Time: O(n)
// Space: O(2n)

package main

func longestValidParentheses(s string) int {
	stack := make([]int, 0, len(s))
	rec := make([]int, len(s))
	max := 0
	for i := 0; i < len(s)-max; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				// pop
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				k := i - j + 1
				if i-k > 0 {
					k += rec[i-k]
				}
				if k > max {
					max = k
				}
				rec[i] = k
			}
		}
	}
	return max
}
