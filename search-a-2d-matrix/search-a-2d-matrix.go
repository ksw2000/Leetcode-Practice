// https://leetcode.com/problems/search-a-2d-matrix/
// 0 ms 2.6 MB

package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	start := 0
	end := rows - 1

	if target < matrix[start][0] || target > matrix[end][cols-1] {
		return false
	}

	var middle int
	for start <= end {
		middle = (start + end) >> 1
		if matrix[middle][0] > target {
			end = middle - 1
		} else if matrix[middle][0] < target {
			start = middle + 1
		} else {
			return true
		}
	}

	row := end
	start = 0
	end = cols - 1
	for start <= end {
		middle = (start + end) >> 1
		if matrix[row][middle] > target {
			end = middle - 1
		} else if matrix[row][middle] < target {
			start = middle + 1
		} else {
			return true
		}
	}
	return false
}
