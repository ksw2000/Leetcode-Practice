// https://leetcode.com/problems/spiral-matrix-ii/
// 0 ms 2.1 MB

package main

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	n-- // loop counter
	r1 := 0
	c1 := 0
	r3 := n
	c3 := n
	j1 := 1
	j2 := n + 1
	j3 := n + j2
	j4 := n + j3

	for n > 0 {
		for i := 0; i < n; i++ {
			ret[r1][c1] = j1
			ret[c1][r3] = j2
			ret[r3][c3] = j3
			ret[c3][r1] = j4
			c1++
			j1++
			c3--
			j2++
			j3++
			j4++
		}

		r1++
		r3--
		n--
		c1 = c1 - n
		c3 = c3 + n

		n--
		j1 = j4
		j2 = j1 + n
		j3 = j2 + n
		j4 = j3 + n
	}
	if n&1 == 0 {
		ret[r1][c1] = j1
	}
	return ret
}

// 0 ms 2 MB
func generateMatrixAnotherWay(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	steps := n - 1
	r := 0
	c := 0
	j := 1
	for steps > 0 {
		for i := 0; i < steps; i++ {
			ret[r][c] = j
			c++
			j++
		}
		for i := 0; i < steps; i++ {
			ret[r][c] = j
			r++
			j++
		}
		for i := 0; i < steps; i++ {
			ret[r][c] = j
			c--
			j++
		}
		for i := 0; i < steps; i++ {
			ret[r][c] = j
			r--
			j++
		}
		steps -= 2
		r++
		c++
	}
	if n&1 == 1 {
		ret[r][c] = j
	}
	return ret
}
