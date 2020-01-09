// URL : https://leetcode.com/problems/divide-two-integers/
// 0 ms	2.4 MB
package main
import "fmt"
func divide(dividend int, divisor int) int {
    var tmp int64
    tmp = int64(dividend) / int64(divisor)
    if(tmp > 2147483647){
        return 2147483647
    }else if(tmp < -2147483648){
        return -2147483648
    }
    return int(tmp)
}
func main(){
    fmt.Printf("divide(%d, %d) = %d\n", 10, 3, divide(10, 3));
    fmt.Printf("divide(%d, %d) = %d\n", 7, -3, divide(7, -3));
}
