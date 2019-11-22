// URL : https://leetcode.com/problems/search-insert-position/
//4 ms	3.1 MB
package main
import "fmt"
func searchInsert(nums []int, target int) int {
    for i:=0; i<len(nums); i++{
        if nums[i]>= target{
            return i
        }
    }
    return len(nums)
}

func main(){
    arr:=[]int{1,3,5,6}
    fmt.Println(searchInsert(arr,5))
    fmt.Println(searchInsert(arr,7))
    fmt.Println(searchInsert(arr,0))
    arr2:=[]int{}
    fmt.Println(searchInsert(arr2,10))
}
