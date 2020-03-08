// URL : https://leetcode.com/problems/merge-k-sorted-lists/
// 8 ms	5.7 MB
package main
import "fmt"

type ListNode struct{
    Val int
    Next *ListNode
}

func printLinkedList(root *ListNode){
    for current:=root; current!=nil; current = current.Next{
        fmt.Printf("%d -> ", current.Val)
    }
    fmt.Printf("nil\n")
}

func appendList(root **ListNode, node *ListNode, val int) *ListNode{
    new_node := new(ListNode)
    new_node.Val = val
    new_node.Next = nil

    if *root == nil{
        *root = new_node
        node = *root
    }else{
        node.Next = new_node
        node = node.Next
    }

    return node
}

//----------------------------------- Answer Begin -----------------------------------//
func mergeTwo(list1 *ListNode, list2 *ListNode) *ListNode{
    if list1 == nil{
        list2, list1 = list1, list2
    }

    previous := new(ListNode)
    previous.Next = list1
    head1 := previous

    for list1!=nil && list2!=nil{
        if list1.Val > list2.Val{
            temp := list2.Next
            previous.Next = list2
            list2.Next = list1
            previous = previous.Next
            list2 = temp
        }else{
            previous = list1
            list1 = list1.Next
        }
    }

    if list2!=nil{
        previous.Next = list2
    }

    return head1.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0{
        return nil
    }

    for len(lists)>1{
        l1 := lists[0]
        l2 := lists[1]
        lists = lists[2:]
        lists = append(lists, mergeTwo(l1,l2))
    }

    return lists[0]
}
//----------------------------------- Answer END -----------------------------------//

func main(){
    var head, tail *ListNode
    head = nil
    tail = head
    tail = appendList(&head, tail, 1)
    tail = appendList(&head, tail, 4)
    tail = appendList(&head, tail, 5)

    var head2, tail2 *ListNode
    head2 = nil
    tail2 = head2
    tail2 = appendList(&head2, tail2, 1)
    tail2 = appendList(&head2, tail2, 3)
    tail2 = appendList(&head2, tail2, 4)

    var head3, tail3 *ListNode
    head3 = nil
    tail3 = head3
    tail3 = appendList(&head3, tail3, 2)
    tail3 = appendList(&head3, tail3, 6)

    listArray := []*ListNode{head, head2, head3}

    printLinkedList(mergeKLists(listArray))
}
