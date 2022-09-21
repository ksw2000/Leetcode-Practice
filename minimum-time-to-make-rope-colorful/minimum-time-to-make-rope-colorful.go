// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
package main

func minCost(colors string, neededTime []int) int {
	cost := 0
	for i := 1; i < len(colors); i++ {
		if colors[i] == colors[i-1] {
			// keep the max at the left
			if neededTime[i] < neededTime[i-1] {
				cost += neededTime[i]
				neededTime[i] = neededTime[i-1]
			} else {
				cost += neededTime[i-1]
			}
		}
	}
	return cost
}
