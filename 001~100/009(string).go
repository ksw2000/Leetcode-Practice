// URL : https://leetcode.com/problems/palindrome-number/
package main
import "fmt"
func isPalindrome(x int) bool{
    if x<0{
        return false
    }
    var num string
    num=fmt.Sprintf("%d",x)

    half:=len(num)>>1
    last:=len(num)-1
    for i:=0; i<half; i++{
        if num[i] != num[last-i]{
            return false
        }
    }
    return true
}
func main(){
    fmt.Println(isPalindrome(-28))
    fmt.Println(isPalindrome(8987))
    fmt.Println(isPalindrome(88))
    fmt.Println(isPalindrome(878))
    fmt.Println(isPalindrome(8998))
    fmt.Println(isPalindrome(1102))
}
