// URL : https://leetcode.com/problems/palindrome-number/
// 16 ms 5.2 MB
// 1. Reverse the number
// 2. Compare
package main
import "fmt"
func isPalindrome(x int) bool {
    if x < 0{
        return false
    }

    if x < 10{
        return true
    }

    tmp := x
    reverse := 0
    for tmp > 0{
        reverse = reverse * 10 + tmp % 10
        tmp /= 10
    }

    return x == reverse
}

func main(){
    fmt.Println(isPalindrome(-28))
    fmt.Println(isPalindrome(8987))
    fmt.Println(isPalindrome(88))
    fmt.Println(isPalindrome(878))
    fmt.Println(isPalindrome(8998))
    fmt.Println(isPalindrome(1102))
}
