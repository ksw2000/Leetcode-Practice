// https://leetcode.com/problems/insertion-sort-list/
// 6 ms 3.2MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	pseudoHead := &ListNode{
		Next: head,
	}

	if head == nil {
		return head
	}

	for cur := head; cur.Next != nil; {
		if cur.Next.Val >= cur.Val {
			cur = cur.Next
			continue
		}
		// pre   pre.next ... cur  cur.next
		//     ↑                         |
		//     |_________________________|

		pre := pseudoHead
		for pre.Next.Val < cur.Next.Val {
			pre = pre.Next
		}
		cur.Next.Next, pre.Next, cur.Next = pre.Next, cur.Next, cur.Next.Next
	}

	return pseudoHead.Next
}
