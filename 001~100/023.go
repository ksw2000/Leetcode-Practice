// URL : https://leetcode.com/problems/merge-k-sorted-lists/
// 8 ms	5.6 MB
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
func merge(i int, j int, lists[]*ListNode){
    if i < j{
        mid := (i+j)>>1
        merge(i, mid, lists)
        merge(mid+1, j, lists)

        list1, list2 := lists[i], lists[j]
        if lists[i] == nil{
            list2, list1 = lists[i], lists[j]
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

        head1 = head1.Next
        for k:=i; k<=j; k++{
            lists[k] = head1
        }
    }
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists)>0{
        merge(0, len(lists)-1, lists)
        return lists[0]
    }
    return nil
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
