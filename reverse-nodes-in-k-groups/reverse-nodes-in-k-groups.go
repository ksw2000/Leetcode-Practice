// https://leetcode.com/problems/reverse-nodes-in-k-group/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(x *ListNode, y *ListNode) {
	if x.Next == y {
		tmp := x.Next
		x.Next = y.Next
		y.Next = x.Next.Next
		x.Next.Next = tmp
	} else {
		tmp := x.Next.Next
		x.Next.Next = y.Next.Next
		y.Next.Next = tmp
		tmp = x.Next
		x.Next = y.Next
		y.Next = tmp
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	redundantHead := &ListNode{
		Next: head,
	}

	buffer := make([]*ListNode, k)
	i := 0
	current := redundantHead
	for current.Next != nil {
		buffer[i] = current
		i++
		if i == k {
			current = buffer[0].Next
			// swap
			for j := 0; j < k/2; j++ {
				swapNodes(buffer[j], buffer[k-1-j])
				buffer[j+1] = buffer[j].Next
			}
			i = 0
		} else {
			current = current.Next
		}
	}
	return redundantHead.Next
}
