// URL : https://leetcode.com/problems/container-with-most-water/
// 12 ms 5.8 MB
package main
import "fmt"

//----------------------------------- Answer Begin -----------------------------------//
func min(a int,b int)int{
    if a>b{
        return b
    }
    return a
}

func maxArea(height []int) int {
    max := 0
    i:=0
    j:=len(height)-1
    for i<j{
        if min(height[i], height[j])*(j-i) > max{
            max = min(height[i], height[j])*(j-i)
        }

        if height[i]>height[j]{
            j--
        }else{
            i++
        }
    }
    return max
}
//----------------------------------- Answer END -----------------------------------//

func main(){
    arr := []int{1,8,6,2,5,4,8,3,7}
    fmt.Println(maxArea(arr))
}
