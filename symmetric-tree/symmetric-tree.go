// https://leetcode.com/problems/symmetric-tree/
// 4 ms 2.8 MB

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTwoNodes(root.Left, root.Right)
}

func isSymmetricTwoNodes(m, n *TreeNode) bool {
	if m == nil || n == nil {
		return m == n
	}
	if m.Val == n.Val {
		return isSymmetricTwoNodes(m.Left, n.Right) && isSymmetricTwoNodes(m.Right, n.Left)
	}
	return false
}
