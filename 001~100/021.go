// URL : https://leetcode.com/problems/merge-two-sorted-lists/
// 0 ms	2.6 MB
package main
import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

//----------------------------------- Answer Begin -----------------------------------//
func compare(a,b int)int{
    if a>b{
        return 1
    }else if a==b{
        return 0
    }
    return -1
}

func linkAdd(head **ListNode, end **ListNode, val int){
    newList:=new(ListNode)
    (*newList).Val=val
    (*newList).Next=nil

    if (*head)==nil && (*end)==nil{
        (*head)=newList
        (*end)=newList
    }else{
        (*(*end)).Next=newList
        (*end)=newList
    }
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var current1, current2, head_l3, l3 *ListNode
    current1=l1
    current2=l2
    for current1!=nil && current2!=nil{
        switch compare((*current1).Val, (*current2).Val) {
        case -1:
            linkAdd(&head_l3, &l3, (*current1).Val)
            current1=(*current1).Next
        case 0:
            linkAdd(&head_l3, &l3, (*current1).Val)
            linkAdd(&head_l3, &l3, (*current2).Val)
            current1=(*current1).Next
            current2=(*current2).Next
        case 1:
            linkAdd(&head_l3, &l3, (*current2).Val)
            current2=(*current2).Next
        }
    }

    for current1!=nil{
        linkAdd(&head_l3, &l3, (*current1).Val)
        current1=(*current1).Next
    }
    for current2!=nil{
        linkAdd(&head_l3, &l3, (*current2).Val)
        current2=(*current2).Next
    }
    return head_l3
}
//----------------------------------- Answer END -----------------------------------//

func main(){
    var head_l1, l1, head_l2, l2 *ListNode
    linkAdd(&head_l1, &l1, 1)
    linkAdd(&head_l1, &l1, 2)
    linkAdd(&head_l1, &l1, 4)

    linkAdd(&head_l2, &l2, 1)
    linkAdd(&head_l2, &l2, 3)
    linkAdd(&head_l2, &l2, 4)

    tmp:=mergeTwoLists(head_l1, head_l2)
    for current:=tmp; current!=nil; current=(*current).Next{
        fmt.Print((*current).Val)
        fmt.Print(" ")
    }
}
