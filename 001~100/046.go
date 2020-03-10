// URL : https://leetcode.com/problems/permutations/
// 0 ms	2.7 MB
package main
import "fmt"

func shift(list []int) []int{
    temp := list[0]
    list = list[1:]
    list = append(list, temp)
    return list
}

func next(fixedIndex int, list []int, answer *[][]int){
    if fixedIndex < len(list)-1{
        for i := 0; i<len(list)-1-fixedIndex; i++{
            l := shift(list[fixedIndex+1:])
            next(fixedIndex+1, append(list[:fixedIndex+1], l...), answer)
        }
    }else{
        new_list := make([]int, len(list))
        copy(new_list, list)
        *answer = append(*answer, new_list)
    }
}

func permute(nums []int) [][]int {
    answer := [][]int{}
    next(-1, nums, &answer)
    return answer
}

func main(){
    fmt.Println(permute([]int{}))
    fmt.Println(permute([]int{1}))
    fmt.Println(permute([]int{1,2}))
    fmt.Println(permute([]int{1,2,3}))
    fmt.Println(permute([]int{1,2,3,4}))
}
