// URL : https://leetcode.com/problems/powx-n/
// 4268 ms	2.1 MB
package main
import "fmt"

func myPow(x float64, n int) float64 {
    var tmp float64
    var m int

    if x==0{
        return 0
    }else if x==1{
        return 1
    }else if x==-1{
        if n%2==1{
            return -1
        }else{
            return 1;
        }
    }

    if n>=0 {
        m = n
    }else{
        m = n*(-1)
    }

    tmp = 1
    for i:=0; i<m; i++{
        tmp *= x
    }

    if m!=n{
        return 1 / tmp
    }

    return tmp
}

func main(){
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, 10, myPow(2.00000,10))
    fmt.Printf("pow(%f, %d) = %f\n", 2.10000, 3, myPow(2.10000, 3))
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, -2, myPow(2.00000, -2))
    fmt.Printf("pow(%f, %d) = %f\n", 2.00000, 0, myPow(2.00000, 0))
    fmt.Printf("pow(%f, %d) = %f\n", 1.00000, -2147483648, myPow(1.00000, -2147483648))
}
