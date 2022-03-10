// https://leetcode.com/problems/rotate-list/
// 2 ms 2.6 MB

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	end := head
	for i := 0; i < k; i++ {
		if end.Next == nil {
			end = head
			k = k % (i + 1)
			for j := 0; j < k; j++ {
				end = end.Next
			}
			break
		} else {
			end = end.Next
		}
	}
	start := head
	for end.Next != nil {
		end = end.Next
		start = start.Next
	}
	end.Next = head
	ret := start.Next
	start.Next = nil
	return ret
}
