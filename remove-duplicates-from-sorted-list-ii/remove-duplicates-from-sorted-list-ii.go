// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// 4 ms	3 MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	pseudoHead := &ListNode{
		Next: head,
	}
	current := pseudoHead
	for visit := head; visit != nil; visit = visit.Next {
		if visit.Next != nil && visit.Val == visit.Next.Val {
			for visit.Next != nil && visit.Val == visit.Next.Val {
				visit = visit.Next
			}
		} else {
			current.Next = visit
			current = current.Next
		}
	}
	if current != nil {
		current.Next = nil
	}
	return pseudoHead.Next
}
