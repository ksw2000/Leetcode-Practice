// https://leetcode.com/problems/merge-k-sorted-lists/submissions/884241337/
// 7 ms 5.5 MB

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwo(list1, list2 *ListNode) *ListNode {
	start := &ListNode{}
	end := start
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			end.Next = list1
			list1 = list1.Next
		} else {
			end.Next = list2
			list2 = list2.Next
		}
		end = end.Next
	}
	for list1 != nil {
		end.Next = list1
		end = end.Next
		list1 = list1.Next
	}
	for list2 != nil {
		end.Next = list2
		end = end.Next
		list2 = list2.Next
	}
	return start.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		lists = append(lists, mergeTwo(lists[0], lists[1]))
		lists = lists[2:]
	}

	return lists[0]
}
