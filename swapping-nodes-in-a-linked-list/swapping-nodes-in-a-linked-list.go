// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
// 148 ms 9.1MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	traverserHead := head
	traverserTail := head
	for i := 0; i < k-1; i++ {
		traverserTail = traverserTail.Next
	}

	swapping1 := traverserTail
	traverserTail = traverserTail.Next
	for traverserTail != nil {
		traverserTail = traverserTail.Next
		traverserHead = traverserHead.Next
	}
	swapping2 := traverserHead
	if swapping1 != swapping2 {
		swapping1.Val = swapping1.Val + swapping2.Val
		swapping2.Val = swapping1.Val - swapping2.Val
		swapping1.Val = swapping1.Val - swapping2.Val
	}

	return head
}
