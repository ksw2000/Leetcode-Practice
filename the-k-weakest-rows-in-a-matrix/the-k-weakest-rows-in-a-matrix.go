// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/submissions/
// 8 ms	4.9 MB

package main

func kWeakestRows(mat [][]int, k int) []int {
	rows := len(mat)
	cols := len(mat[0])
	visited := make([]bool, rows)
	ret := make([]int, 0, k)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if mat[j][i] == 0 && !visited[j] {
				ret = append(ret, j)
				visited[j] = true
				k--
			}
			if k == 0 {
				return ret
			}
		}
	}
	// append the remaining rows if it has not been visited
	for i := 0; k > 0; i++ {
		if !visited[i] {
			ret = append(ret, i)
			k--
		}
	}
	return ret
}
