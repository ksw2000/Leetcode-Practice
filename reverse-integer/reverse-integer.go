// URL : https://leetcode.com/problems/reverse-integer/
// 0 ms 2.3 MB
package main
import "fmt"
func reverse(x int) int {
    y := 0
    for ; x != 0; x /= 10 {
        y = y * 10 + x % 10
    }

    if y > 2147483647 || y < -2147483648{
        return 0
    }
    return y
}

func main(){
    fmt.Println(reverse(1234))
    fmt.Println(reverse(12345))
    fmt.Println(reverse(-12345))
    fmt.Println(reverse(2147483647))
    fmt.Println(reverse(1534236469))
    fmt.Println(reverse(-2147483648))
}
