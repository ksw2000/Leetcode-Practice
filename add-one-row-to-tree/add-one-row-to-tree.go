// https://leetcode.com/problems/add-one-row-to-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	queue := []*TreeNode{root}
	depth -= 2
	for depth > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[i]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		queue = queue[size:]
		depth--
	}
	for _, node := range queue {
		node.Left = &TreeNode{
			Val:  val,
			Left: node.Left,
		}
		node.Right = &TreeNode{
			Val:   val,
			Right: node.Right,
		}
	}
	return root
}
