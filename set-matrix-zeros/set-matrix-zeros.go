// https://leetcode.com/problems/set-matrix-zeroes/
// 12 ms 6.3 MB

package main

func setZeroes(matrix [][]int) {
	var setFirstRowZeros, setFirstColumnZeros bool
	r := len(matrix)
	c := len(matrix[0])
	for i := 0; i < r && !setFirstColumnZeros; i++ {
		setFirstColumnZeros = matrix[i][0] == 0
	}
	for i := 0; i < c && !setFirstRowZeros; i++ {
		setFirstRowZeros = matrix[0][i] == 0
	}
	// if  (u, v) is zero
	// set (u, 0) zero
	// set (0, v) zero
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][0]*matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if setFirstRowZeros {
		for i := 0; i < c; i++ {
			matrix[0][i] = 0
		}
	}
	if setFirstColumnZeros {
		for i := 0; i < r; i++ {
			matrix[i][0] = 0
		}
	}
}
