// URL : https://leetcode.com/problems/search-in-rotated-sorted-array/
// 0 ms	2.6 MB
package main
import "fmt"

func search(nums []int, target int) int {
    if len(nums)==0{
        return -1
    }

    start,end := 0, len(nums)-1

    for ;; {
        mid := (start+end)/2
        if(end-start>1){
            if nums[mid] > nums[start]{
                if nums[start] <= target && target <= nums[mid] {
                    start = start
                    end = mid
                    continue
                }else{
                    start = mid+1
                    end = end
                    continue
                }
            }else{
                if nums[mid] <= target && target <= nums[end]{
                    start = mid
                    end = end
                    continue
                }else{
                    start = start
                    end = mid-1
                    continue
                }
            }
        }else{
            if nums[start]==target{
                return start
            }else if nums[end]==target{
                return end
            }
        }

        return -1
    }
}

func main(){
    arr1:=[]int{4,5,6,7,0,1,2}
    fmt.Println(search(arr1,7))

    arr2:=[]int{4,5,6,7,0,1,2}
    fmt.Println(search(arr2,3))

    arr3:=[]int{}
    fmt.Println(search(arr3,3))

    arr4:=[]int{1,3}
    fmt.Println(search(arr4,3))

    arr5:=[]int{1,3,5}
    fmt.Println(search(arr5,1))
}
