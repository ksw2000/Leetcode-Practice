// https://leetcode.com/problems/01-matrix/description/
// 49 ms 7.4 MB
package main

func updateMatrix(mat [][]int) [][]int {
	solved := make([]int, 0, len(mat)*len(mat[0]))
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] != 0 {
				mat[i][j] = -1
			} else {
				solved = append(solved, i*len(mat[0])+j)
			}
		}
	}

	for k := 0; k < len(solved); k++ {
		i := solved[k] / len(mat[0])
		j := solved[k] % len(mat[0])
		if i-1 >= 0 && mat[i-1][j] == -1 {
			mat[i-1][j] = mat[i][j] + 1
			solved = append(solved, solved[k]-len(mat[0]))
		}

		if i+1 < len(mat) && mat[i+1][j] == -1 {
			mat[i+1][j] = mat[i][j] + 1
			solved = append(solved, solved[k]+len(mat[0]))
		}

		if j-1 >= 0 && mat[i][j-1] == -1 {
			mat[i][j-1] = mat[i][j] + 1
			solved = append(solved, solved[k]-1)
		}

		if j+1 < len(mat[0]) && mat[i][j+1] == -1 {
			mat[i][j+1] = mat[i][j] + 1
			solved = append(solved, solved[k]+1)
		}
	}
	return mat
}
