// https://leetcode.com/problems/generate-parentheses/
// 0 ms 2.8 MB
package main

func generate(left int, right int, n int, result string, answer *[]string) {
	if left+right == n<<1 {
		*answer = append(*answer, result)
	}
	if left < n {
		generate(left+1, right, n, result+"(", answer)
	}
	if left > right {
		generate(left, right+1, n, result+")", answer)
	}
}

func generateParenthesis(n int) []string {
	answer := []string{}
	generate(1, 0, n, "(", &answer)
	return answer
}
