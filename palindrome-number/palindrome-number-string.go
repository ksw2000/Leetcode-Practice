// URL : https://leetcode.com/problems/palindrome-number/
// 8 ms 5.5 MB
// 1. Convert number to string
// 2. Compare
package main
import "fmt"
func isPalindrome(x int) bool{
    if x < 0{
        return false
    }

    num  := fmt.Sprintf("%d", x)

    for i := 0; i < (len(num) >> 1); i++ {
        if num[i] != num[len(num) - 1 - i]{
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
