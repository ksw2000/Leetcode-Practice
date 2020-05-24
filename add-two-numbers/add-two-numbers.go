// https://leetcode.com/problems/add-two-numbers/
// 8 ms	5 MB
package main
import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct{
    Val int
    Next *ListNode
}

//----------------------------------- Answer Begin -----------------------------------//
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    root  := new(ListNode)
    tail  := root
    carry := 0
    sum   := 0

    for l1!=nil && l2!=nil {
        sum       = l1.Val + l2.Val + carry
        carry     = sum / 10
        tail.Next = new(ListNode)
        tail      = tail.Next
        tail.Val  = sum % 10
        l1        = l1.Next
        l2        = l2.Next
    }

    for l1!=nil {
        sum       = l1.Val + carry
        carry     = sum / 10
        tail.Next = new(ListNode)
        tail      = tail.Next
        tail.Val  = sum % 10
        l1        = l1.Next
    }

    for l2!=nil {
        sum       = l2.Val + carry
        carry     = sum / 10
        tail.Next = new(ListNode)
        tail      = tail.Next
        tail.Val  = sum % 10
        l2        = l2.Next
    }

    if carry > 0 {
        tail.Next = new(ListNode)
        tail      = tail.Next
        tail.Val  = carry
    }

    return root.Next
}
//----------------------------------- Answer END -----------------------------------//
func makeLL(list []int) *ListNode{
    var root *ListNode
    for i := len(list)-1; i>=0; i--{
        n     := new(ListNode)
        n.Val  = list[i]
        n.Next = root
        root   = n
    }
    return root
}

func printLL(l *ListNode){
    for cur := l; cur != nil; cur = cur.Next{
        fmt.Printf("%d -> ", cur.Val)
    }
    fmt.Printf("nil\n")
}

func main(){
    l1 := makeLL([]int{2, 4, 3})
    l2 := makeLL([]int{5, 6, 4})
    l3 := addTwoNumbers(l1, l2)
    printLL(l1)
    printLL(l2)
    printLL(l3)
}
