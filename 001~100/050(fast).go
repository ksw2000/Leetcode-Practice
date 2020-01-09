// URL : https://leetcode.com/problems/powx-n/
// 0 ms	2.1 MB
package main
import "fmt"

func myPow(x float64, n int) float64 {
    if(n!=1){
        if(n==0){
            return 1
        }else if(n<0){
            return 1 / myPow(x, -n)
        }else if(n>0){
            if n%2==0 {
                return myPow(x*x, n >> 1)
            }else{
                return x * myPow(x, n-1)
            }
        }
    }

    return x
}

func main(){
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, 10, myPow(2.00000,10))
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, 9, myPow(2.00000,9))
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, 8, myPow(2.00000,8))
    fmt.Printf("pow(%f, %d) = %f\n", 2.10000, 3, myPow(2.10000, 3))
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, -2, myPow(2.00000, -2))
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, 0, myPow(2.00000, 0))
    fmt.Printf("pow(%f, %d) = %f\n", 1.00000, -2147483648, myPow(1.00000, -2147483648))
}
