// https://leetcode.com/problems/copy-list-with-random-pointer/
// 0 ms 3.6 MB
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mapping := map[*Node]int{}
	newList := [1000]Node{}
	i := 0
	for current := head; current != nil; current = current.Next {
		newList[i] = Node{
			Val:  current.Val,
			Next: &newList[i+1],
		}
		mapping[current] = i
		i++
	}
	if i > 0 {
		newList[i-1].Next = nil
	} else {
		return nil
	}
	i = 0
	for current := head; current != nil; current = current.Next {
		if idx, ok := mapping[current.Random]; ok {
			newList[i].Random = &newList[idx]
		} else {
			newList[i].Random = nil
		}
		i++
	}

	return &newList[0]
}
