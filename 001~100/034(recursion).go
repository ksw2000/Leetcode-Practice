// URL : https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 8 ms	4.1 MB
package main
import "fmt"

func findLeft(start int, end int, target int, nums *[]int) int{
    if start < end{
        mid := (start + end) >> 1
        if (*nums)[mid] >= target{
            return findLeft(start, mid, target, nums)
        }else{
            return findLeft(mid+1, end, target, nums)
        }
    }else{
        //fmt.Printf("start:%d end:%d \n",start,end)
        if (*nums)[start] == target{
            return start
        }
    }
    return -1
}

func findRight(start int, end int, target int, nums *[]int) int{
    if start < end{
        mid := (start + end) >> 1 + 1
        if (*nums)[mid] <= target{
            return findRight(mid, end, target, nums)
        }else{
            return findRight(start, mid-1, target, nums)
        }
    }else{
        //fmt.Printf("start:%d end:%d \n",start,end)
        if (*nums)[end] == target{
            return end
        }
    }
    return -1
}

func searchRange(nums []int, target int) []int {
    if len(nums)>0{
        if left := findLeft(0, len(nums)-1, target, &nums); left != -1{
            return []int{left, findRight(left, len(nums)-1, target, &nums)}
        }
    }
    return []int{-1, -1}
}

func main(){
    fmt.Println(searchRange([]int{1,2,3}, 2))
    fmt.Println(searchRange([]int{1,2,3,5}, 3))
    fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
    fmt.Println(searchRange([]int{8,8,8,8,8,8}, 8))
    fmt.Println(searchRange([]int{8,8,8,8,8,8,8}, 8))
    fmt.Println(searchRange([]int{3,3,3}, 3))
    fmt.Println(searchRange([]int{1}, 1))
    fmt.Println(searchRange([]int{3,5,6,7}, 4))
    fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
    fmt.Println(searchRange([]int{}, 0))
}
