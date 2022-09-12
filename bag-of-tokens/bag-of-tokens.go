// https://leetcode.com/problems/bag-of-tokens/
// 5 ms	3 MB

package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	score := 0
	i, j := 0, len(tokens)
	for i < j {
		if power >= tokens[i] {
			score++
		} else if score > 0 && i < j-1 {
			power += tokens[j-1]
			j--
		} else {
			break
		}
		power -= tokens[i]
		i++
	}
	return score
}
