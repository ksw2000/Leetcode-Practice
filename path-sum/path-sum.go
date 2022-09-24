// https://leetcode.com/problems/path-sum/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return root != nil && ((root.Val == targetSum && root.Left == nil && root.Right == nil) || hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val))
}

func hasPathSumWithGoroutine(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	hasR := make(chan bool)
	go func(hasR chan bool, target int) {
		hasR <- hasPathSum(root.Right, target)
	}(hasR, targetSum-root.Val)
	hasL := hasPathSum(root.Left, targetSum-root.Val)

	return hasL || <-hasR
}
