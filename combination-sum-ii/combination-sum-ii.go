// https://leetcode.com/problems/combination-sum-ii/
// 0 ms 3.2 MB

package main

type Node struct {
	sum      int
	previous *Node
}

type Num struct {
	num     int
	counter int
}

func combinationSum2(candidates []int, target int) [][]int {
	counter := make([]int, 51)
	for _, c := range candidates {
		counter[c]++
	}

	numbers := []Num{}
	for i := len(counter) - 1; i >= 0; i-- {
		if counter[i] != 0 {
			numbers = append(numbers, Num{
				num:     i,
				counter: counter[i],
			})
		}
	}

	root := &Node{
		sum:      0,
		previous: nil,
	}
	queue := []*Node{root}

	res := [][]int{}
	for _, n := range numbers {
		size := len(queue)
		// dequeue
		for j := 0; j < size; j++ {
			current := queue[j]
			sum := current.sum
			queue = append(queue, current)
			for i := 0; i < n.counter; i++ {
				sum += n.num
				if sum < target {
					// Add to the end of queue
					current = &Node{
						sum:      sum,
						previous: current,
					}
					queue = append(queue, current)
				} else if sum == target {
					row := []int{n.num}
					for c := current; c != root; c = c.previous {
						row = append(row, c.sum-c.previous.sum)
					}
					res = append(res, row)
				} else {
					break
				}
			}
		}
		// remove the traversed items
		queue = queue[size:]
	}
	return res
}
