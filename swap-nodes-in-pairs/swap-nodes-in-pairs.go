// https://leetcode.com/problems/swap-nodes-in-pairs/
// 0 ms 2.1 MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{
		Next: head,
	}

	for p1 := head; p1 != nil && p1.Next != nil && p1.Next.Next != nil; {
		p1.Next, p1.Next.Next, p1.Next.Next.Next = p1.Next.Next, p1.Next.Next.Next, p1.Next
		p1 = p1.Next.Next
	}

	return head.Next
}
