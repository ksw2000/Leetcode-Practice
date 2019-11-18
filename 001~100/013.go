// URL : https://leetcode.com/problems/roman-to-integer/
//4 ms	3 MB
package main
import "fmt"

func romanToInt(s string) int {
    var precedence = map[byte]int{
        'I' : 1,
        'V' : 5,
        'X' : 10,
        'L' : 50,
        'C' : 100,
        'D' : 500,
        'M' : 1000,
    }
    len_minus_one := len(s)-1
    sum := 0
    i := 0
    for i < len_minus_one {
        precedence_i := precedence[s[i]]
        precedence_i_next := precedence[s[i+1]]
        if precedence_i >= precedence_i_next{
            sum += precedence_i
        }else{
            sum -= precedence_i
        }
        i+=1
    }
    return sum+precedence[s[i]]
}

func main(){
    fmt.Println(romanToInt("MCMXCIV"))
}
