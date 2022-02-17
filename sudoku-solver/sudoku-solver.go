// https://leetcode.com/problems/sudoku-solver/
// 0 ms 2.4 MB

package main

import (
	"fmt"
)

func subtaskSolveSudoku(board [][]byte) (done bool) {
	unsolved := []int8{}
	row := [9]uint16{}
	col := [9]uint16{}
	box := [9]uint16{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			k := board[i][j]
			if k == '.' {
				unsolved = append(unsolved, int8(i*9+j))
			} else {
				row[i] |= 1 << int(k-'1')
				col[j] |= 1 << int(k-'1')
				box[(i/3)*3+j/3] |= 1 << int(k-'1')
			}
		}
	}

	continueEnqueue := 0
	for len(unsolved) != 0 {
		// deQueue
		pos := unsolved[0]
		unsolved = unsolved[1:]

		i, j := pos/9, pos%9
		candidates := []int{}
		mask := row[i] | col[j] | box[(i/3)*3+j/3]
		for k := 0; k < 9; k++ {
			if mask&1 == 0 {
				candidates = append(candidates, k)
			}
			mask = mask >> 1
		}

		if len(candidates) == 0 {
			return false
		} else if len(candidates) == 1 {
			board[i][j] = byte(candidates[0] + '1')
			row[i] |= 1 << int(candidates[0])
			col[j] |= 1 << int(candidates[0])
			box[(i/3)*3+j/3] |= 1 << int(candidates[0])
			continueEnqueue = 0
		} else {
			continueEnqueue++
			if continueEnqueue == len(unsolved) {
				for _, candidate := range candidates {
					board[i][j] = byte(candidate + '1')
					if !subtaskSolveSudoku(board) {
						// undo
						for _, pos := range unsolved {
							board[pos/9][pos%9] = byte('.')
						}
						continue
					}
					return true
				}
				return false
			}
			unsolved = append(unsolved, pos)
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	subtaskSolveSudoku(board)
}

func main() {
	input := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'}}
	solveSudoku(input)
	showBoard(input)
}

func showBoard(board [][]byte) {
	for _, i := range board {
		for _, j := range i {
			fmt.Printf("%c ", j)
		}
		fmt.Println()
	}
}
