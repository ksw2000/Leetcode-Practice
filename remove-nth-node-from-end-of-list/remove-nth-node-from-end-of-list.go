// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// 2 ms 2.2 MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pseudoHead := &ListNode{
		Next: head,
	}
	left := pseudoHead
	for i := 0; i < n+1; i++ {
		left = left.Next
	}
	right := pseudoHead
	for left != nil {
		left = left.Next
		right = right.Next
	}
	right.Next = right.Next.Next
	return pseudoHead.Next
}
