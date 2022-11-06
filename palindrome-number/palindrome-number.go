// URL : https://leetcode.com/problems/palindrome-number/
// 4 ms	5.2 MB
// 1. Do not transfer number to string
//    But store in a stack
// 2. Compare
package main
import "fmt"
func isPalindrome(x int) bool{
    if x < 0{
        return false
    }

    stack := [10]int{}
    index := -1
    for x > 0{
        index++
        stack[index] = x % 10
        x /= 10
    }

    for i:=0; i < ((index + 1) >> 1); i++ {
        if stack[index-i] != stack[i] {
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
