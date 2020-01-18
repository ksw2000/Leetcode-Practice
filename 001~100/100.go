// URL : https://leetcode.com/problems/same-tree/
// 0 ms 2.1 MB
package main
import "fmt"

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

//----------------------------------- Answer Begin -----------------------------------//
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return (p==nil && q==nil) || ((p!=nil && q!=nil && (*p).Val==(*q).Val) && isSameTree((*p).Left, (*q).Left) && isSameTree((*p).Right, (*q).Right))
}
//----------------------------------- Answer END -----------------------------------//

func pushTree(val int, root **TreeNode){
    var tmp *TreeNode = new(TreeNode)
    tmp.Val=val

    var current,previous *TreeNode
    previous = *root

    for current = *root; current != nil; {
        if current.Val > val{
            previous = current
            current = current.Left
        }else{
            previous = current
            current = current.Right
        }
    }

    if *root == nil{
        *root = tmp
    }else if previous.Val > val{
        previous.Left=tmp
    }else{
        previous.Right=tmp
    }
}

func inorder(root *TreeNode){
    if(root!=nil){
        inorder(root.Left)
        fmt.Printf("%d ",root.Val)
        inorder(root.Right)
    }
}

func main(){
    var root,root2,root3 *TreeNode

    var arr  = []int{1,2,3,4,5,6,7}
    var arr2 = []int{1,2,3,4,5,6,7}
    var arr3 = []int{1,2,3,4,5,6}

    for _,val := range arr{
        pushTree(val, &root)
    }

    for _,val := range arr2{
        pushTree(val, &root2)
    }

    for _,val := range arr3{
        pushTree(val, &root3)
    }

    fmt.Println(isSameTree(root,root2));
    fmt.Println(isSameTree(root,root3));
}
