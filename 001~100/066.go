// URL : https://leetcode.com/problems/plus-one/
// 0 ms	2.1 MB
package main
import "fmt"

func plusOne(digits []int) []int {
    last  := len(digits)-1
    carry := 1
    for j:=last; j>=0; j--{
        if digits[j] += carry; digits[j]>=10{
            carry=1
            digits[j] %= 10
        }else{
            carry=0
            return digits
        }
    }

    //overflow
    if carry!=0{
        var newDigits=make([]int,len(digits)+1)
        newDigits[0]=carry
        for j:=0; j<len(digits); j++{
            newDigits[j+1]=digits[j]
        }
        digits = newDigits
    }
    return digits
}

func main(){
    fmt.Println(plusOne([]int{1}))
    fmt.Println(plusOne([]int{9}))
    fmt.Println(plusOne([]int{1,9}))
    fmt.Println(plusOne([]int{9,9}))
    fmt.Println(plusOne([]int{1,2,3}))
}
