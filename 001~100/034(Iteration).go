// URL : https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 8 ms	4.1 MB
package main
import "fmt"

func searchRange(nums []int, target int) []int {
    if len(nums)>0{
        start, end := 0, len(nums)-1

        // Find target
        for start < end{
            mid := (start + end) >> 1
            if nums[mid] > target{
                end = mid
            }else if nums[mid] < target{
                start = mid + 1
            }else{
                nearEnd := end
                // Find left
                for start < end{
                    mid = (start + end) >> 1
                    if nums[mid] >= target{
                        end = mid
                    }else{
                        start = mid + 1
                    }
                }
                left := start

                // Find right
                end = nearEnd
                for start < end{
                    mid = (start + end) >> 1 + 1
                    if nums[mid] <= target{
                        start = mid
                    }else{
                        end = mid-1
                    }
                }

                return []int{left, end}
            }
        }
        if nums[start] == target{
            return []int{start, end}
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
