// URL : https://leetcode.com/problems/merge-sorted-array/
// 0 ms	3.6 MB
package main
import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    index1 := m-1
    index2 := n-1
    index := m+n-1
    for index1>=0 && index2>=0 && index>=0{
        if nums1[index1]>nums2[index2] {
            nums1[index]=nums1[index1]
            index--
            index1--
        }else{
            nums1[index]=nums2[index2]
            index--
            index2--
        }
    }

    for index2>=0 && index>=0{
        nums1[index]=nums2[index2]
        index--
        index2--
    }
}

func main(){
    array1 := []int{1,2,3,0,0,0}
    array2 := []int{2,5,6}
    merge(array1, 3, array2, 3)
    fmt.Println(array1)

    array3 := []int{0}
    array4 := []int{1}
    merge(array3, 0, array4, 1)
    fmt.Println(array3)
}
