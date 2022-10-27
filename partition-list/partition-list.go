// https://leetcode.com/problems/partition-list/
// 0 ms 2.3 MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	pseudoHeadSmall := &ListNode{}
	pseudoHeadBig := &ListNode{}
	small := pseudoHeadSmall
	big := pseudoHeadBig
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			small.Next = cur
			small = cur
		} else {
			big.Next = cur
			big = cur
		}
	}
	big.Next = nil
	small.Next = pseudoHeadBig.Next
	return pseudoHeadSmall.Next
}
