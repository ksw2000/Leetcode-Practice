// https://leetcode.com/problems/two-sum/
// 4 ms	3.8 MB
package main
import "fmt"

func twoSum(nums []int, target int) []int {
    var m = make(map[int]int)
    for k, v := range nums{
        val, ok := m[target-v]
        if ok{
            return []int{val, k}
        }
        m[v] = k
    }
    return []int{}
}

func main(){
    num := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(num, target))
}
