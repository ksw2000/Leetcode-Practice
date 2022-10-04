// https://leetcode.com/problems/binary-tree-preorder-traversal/
// 1 ms 1.9 MB
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(root *TreeNode, order *[]int) {
	if root != nil {
		*order = append(*order, root.Val)
		preorder(root.Left, order)
		preorder(root.Right, order)
	}
}

func preorderTraversal(root *TreeNode) []int {
	order := []int{}
	if root != nil {
		preorder(root, &order)
	}
	return order
}
