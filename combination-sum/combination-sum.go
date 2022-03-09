// https://leetcode.com/problems/combination-sum/
// 0 ms 3.8 MB

package main

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	var ans [][]int
	if len(candidates) > 0 {
		candidate := candidates[len(candidates)-1]
		for count := 0; target >= 0; count++ {
			if target == 0 {
				ans = [][]int{{candidate}}
				for repeat := 1; repeat < count; repeat++ {
					ans[0] = append(ans[0], candidate)
				}
			} else {
				ans = combinationSum(candidates[:len(candidates)-1], target)
				for i := range ans {
					for repeat := 0; repeat < count; repeat++ {
						ans[i] = append(ans[i], candidate)
					}
				}
			}
			// concatenate ret and ans
			ret = append(ret, ans...)
			target -= candidate
		}
	}
	return ret
}
