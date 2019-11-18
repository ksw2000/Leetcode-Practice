// URL : https://leetcode.com/problems/palindrome-number/
package main
import "fmt"
func isPalindrome(x int) bool {
    tmp:=x
    if x<0{
        return false
    }

    inverse:=0
    for tmp>0{
        inverse = inverse*10 + tmp%10
        tmp/=10
    }

    return x==inverse
}
func main(){
    fmt.Println(isPalindrome(-28))
    fmt.Println(isPalindrome(8987))
    fmt.Println(isPalindrome(88))
    fmt.Println(isPalindrome(878))
    fmt.Println(isPalindrome(8998))
    fmt.Println(isPalindrome(1102))
}
