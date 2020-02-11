// URL : https://leetcode.com/problems/container-with-most-water/
// 480 ms 5.7 MB
package main
import "fmt"

func area(h int, g int, side int, max *int){
    if h>g{
        if g*side > *max{
            *max = g*side
        }
    }else{
        if h*side > *max{
            *max = h*side
        }
    }
}

func maxArea(height []int) int {
    max := 0
    for i:=0; i<len(height)-1; i++{
        for j:=i+1; j<len(height); j++{
            area(height[i], height[j], j-i, &max)
        }
    }
    return max
}

func main(){
    arr := []int{1,8,6,2,5,4,8,3,7}
    fmt.Println(maxArea(arr))
}
