// https://leetcode.com/problems/linked-list-cycle/
// 4 ms	4.3 MB
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	turtle := head
	rabbit := head.Next
	for turtle != rabbit && rabbit != nil {
		turtle = turtle.Next
		rabbit = rabbit.Next
		if rabbit != nil {
			rabbit = rabbit.Next
		}
	}
	return turtle == rabbit && turtle != nil
}
