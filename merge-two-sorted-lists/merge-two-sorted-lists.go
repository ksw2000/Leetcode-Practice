// https://leetcode.com/problems/merge-two-sorted-lists/
// 0 ms 2.5 MB
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pseudoHead := &ListNode{}
	end := pseudoHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			end.Next = l1
			l1 = l1.Next
			end = end.Next
		} else if l1.Val > l2.Val {
			end.Next = l2
			l2 = l2.Next
			end = end.Next
		} else {
			end.Next = l1
			l1 = l1.Next
			end = end.Next
			end.Next = l2
			l2 = l2.Next
			end = end.Next
		}
	}

	for l1 != nil {
		end.Next = l1
		l1 = l1.Next
		end = end.Next
	}
	for l2 != nil {
		end.Next = l2
		l2 = l2.Next
		end = end.Next
	}
	return pseudoHead.Next
}
