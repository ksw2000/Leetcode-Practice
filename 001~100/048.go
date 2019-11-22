// URL : https://leetcode.com/problems/rotate-image/
//0 ms	2.7 MB
package main
import "fmt"

//----------------------------------- Answer Begin -----------------------------------//
func rotate(matrix [][]int){
    size := len(matrix)
    for i:=0; i<size/2; i++{
        for j:=i; j<size-i-1; j++{
            matrix[i][j], matrix[j][size-1-i], matrix[size-1-i][size-1-j], matrix[size-1-j][i] = matrix[size-1-j][i], matrix[i][j], matrix[j][size-1-i], matrix[size-1-i][size-1-j]
        }
    }
}
//----------------------------------- Answer END -----------------------------------//

func printM(arr [][]int){
    for i:=0; i<len(arr); i++{
        for j:=0; j<len(arr); j++{
            fmt.Printf("%2d ",arr[i][j])
        }
        fmt.Println()
    }
    fmt.Println()
}

func main(){
    arr:=[][]int{
        {1},
    }
    fmt.Println("before:")
    printM(arr)
    rotate(arr)
    fmt.Println("after:")
    printM(arr)

    arr2:=[][]int{
        {1,2},
        {3,4},
    }
    fmt.Println("before:")
    printM(arr2)
    rotate(arr2)
    fmt.Println("after:")
    printM(arr2)

    arr3:=[][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    fmt.Println("before:")
    printM(arr3)
    rotate(arr3)
    fmt.Println("after:")
    printM(arr3)

    arr4:=[][]int{
        {1,2,3,4},
        {5,6,7,8},
        {9,10,11,12},
        {13,14,15,16},
    }
    fmt.Println("before:")
    printM(arr4)
    rotate(arr4)
    fmt.Println("after:")
    printM(arr4)
}
