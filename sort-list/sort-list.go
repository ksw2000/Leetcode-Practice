// URL : https://leetcode.com/problems/sort-list/
// 36 ms 8.4 MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// l1 and l2 are sorted list.
func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	ret := l3

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}

	if l1 != nil {
		l3.Next = l1
	} else {
		l3.Next = l2
	}
	return ret.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid, cur := head, head.Next
	for cur != nil && cur.Next != nil {
		mid = mid.Next
		cur = cur.Next.Next
	}

	second := mid.Next
	mid.Next = nil
	l1 := sortList(head)
	l2 := sortList(second)

	l3 := &ListNode{}
	head = l3

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}

	if l1 != nil {
		l3.Next = l1
	} else {
		l3.Next = l2
	}

	return head.Next
}
