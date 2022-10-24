// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// 3 ms 3 MB

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	previous := head
	for current := head.Next; current != nil; current = current.Next {
		if current.Val != previous.Val {
			previous.Next = current
			previous = current
		}
	}
	previous.Next = nil
	return head
}
