// https://leetcode.com/problems/add-two-numbers/
// 4 ms	4.7 MB
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
    root  := l1
    carry := 0
    sum   := 0
    var pre *ListNode

    for l1!=nil && l2!=nil {
        sum      = l1.Val + l2.Val + carry
        carry    = sum / 10
        l1.Val   = sum % 10
        pre      = l1
        l1       = l1.Next
        l2       = l2.Next
    }

    // if l2 is too long
    if l2!=nil {
        pre.Next = l2
        l1 = pre.Next
    }

    // if l1 is too long
    for l1!=nil {
        sum       = l1.Val + carry
        carry     = sum / 10
        l1.Val    = sum % 10
        pre       = l1
        l1        = l1.Next
    }


    if carry > 0 {
        pre.Next = new(ListNode)
        pre      = pre.Next
        pre.Val  = carry
    }

    return root
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
    l1 := makeLL([]int{2, 4, 3, 9, 9, 9})
    printLL(l1)
    l2 := makeLL([]int{5, 6, 9, 9, 9})
    printLL(l2)
    l3 := addTwoNumbers(l1, l2)
    printLL(l3)
}
