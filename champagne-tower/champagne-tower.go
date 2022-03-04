// https://leetcode.com/problems/champagne-tower/
// 0 ms	2.1 MB
package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	// because the answer is symmetry
	// champagneTower(?, 10, 7) is equal to (?, 10, 3)
	if query_glass > (query_row >> 1) {
		query_glass = query_row - query_glass
	}

	query_row++
	query_glass++

	// record the exceed champagne
	//        (0,0)
	//     (1,0)(1,1)
	//   (2,0)(2,1)(2,2)
	// (3,0)(3,1)(3,2)(3,3)

	// if we want to find (3,1)
	// we only need to check (0,0), (1,0) (1,1) (2,0) (2,1) (3,1)
	//        (0,0)    /
	//     (1,0)(1,1) /
	//   (2,0)(2,1)  / (2,2)
	// (3,0)(3,1)   / (3,2)(3,3)
	// we need a list which row is 2 (current and previous)
	// and column is query_glass+1 (have added one in advance)

	over := make([]float64, 2*query_glass)
	have := float64(poured) // how much champagne in cup now
	if poured > 1 {
		over[0] = have - 1
		have = 1
	}

	for i := 1; i < query_row; i++ {
		for j := 0; j < query_glass; j++ {
			current := query_glass*(i&1) + j
			previous := query_glass*(i&1^1) + j - 1

			if j == 0 {
				have = over[previous+1] / 2
			} else if j == i {
				have = over[previous] / 2
			} else {
				have = over[previous]/2 + over[previous+1]/2
			}
			if have > 1 {
				over[current] = have - 1
				have = 1
			} else {
				over[current] = 0
			}
		}
	}
	return have
}
