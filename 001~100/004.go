// https://leetcode.com/problems/median-of-two-sorted-arrays/
// 12 ms 5.6 MB
package main
import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    index1 := 0
    index2 := 0
    var median int
    var preMedian int
    var dest int

    count := 0
    dest = (len(nums1)+len(nums2))/2+1

    for index1<len(nums1) && index2<len(nums2) && count < dest{
        preMedian=median
        if(nums1[index1] > nums2[index2]){
            median=nums2[index2]
            index2++
        }else{
            median=nums1[index1]
            index1++
        }
        count++
    }

    for index1<len(nums1) && count < dest{
        preMedian=median
        median=nums1[index1]
        index1++
        count++
    }

    for index2<len(nums2) && count < dest{
        preMedian=median
        median=nums2[index2]
        index2++
        count++
    }

    if (len(nums1)+len(nums2))%2 == 1 {
        return float64(median)
    }
    return (float64(median)+float64(preMedian))/2
}

func main(){
    nums1 := []int{1,3}
    nums2 := []int{2}
    fmt.Println(findMedianSortedArrays(nums1,nums2))
}
