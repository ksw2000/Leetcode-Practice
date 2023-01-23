// https://leetcode.com/problems/longest-path-with-different-adjacent-characters/
// 182 ms 26.6 MB
package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func argmin(a, b int) int {
	if a < b {
		return 0
	}
	return 1
}

func longestPath(parent []int, s string) int {
	haveChildren := make([]bool, len(parent))
	for i := 1; i < len(parent); i++ {
		if s[i] == s[parent[i]] {
			parent[i] = -1
		} else {
			haveChildren[parent[i]] = true
		}
	}

	// add leaves to the queue
	queue := []int{}
	for i := range haveChildren {
		if !haveChildren[i] {
			queue = append(queue, i)
		}
	}

	// find the maximum path from leave for every edges
	// edge[i] -> the link between i and parent[i]

	edge := make([]int, len(parent))

	// initialize
	for p := 0; p < len(queue); p++ {
		if parent[queue[p]] != -1 {
			edge[queue[p]] = 1
		}
	}

	for p := 0; p < len(queue); p++ {
		q := queue[p]
		if parent[q] != -1 && parent[parent[q]] != -1 {
			edge[parent[q]] = max(edge[parent[q]], edge[q]+1)
			queue = append(queue, parent[q])
		}
	}

	//      node0  node1  node2 ...
	// [0]
	// [1]
	table := make([][2]int, len(parent))
	path := 1
	for i := range parent {
		if parent[i] != -1 {
			j := argmin(table[parent[i]][0], table[parent[i]][1])
			table[parent[i]][j] = max(table[parent[i]][j], edge[i])
			if table[parent[i]][0]+table[parent[i]][1]+1 > path {
				path = table[parent[i]][0] + table[parent[i]][1] + 1
			}
		}
	}

	return path
}
