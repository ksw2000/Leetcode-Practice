// URL : https://leetcode.com/problems/integer-to-roman/
//4 ms	3.3 MB
package main
import "fmt"

func intToRoman(num int) string {
    array  := []int   {1,10,100,1000}
    unit   := []string{"I","V","X","L","C","D","M"}
    answer := ""
    for i := len(array)-1; i>=0; i--{
        for num/array[i]>0 {
            switch num/array[i] {
                case 1:
                    answer += unit[i*2]
                case 2:
                    answer += unit[i*2] + unit[i*2]
                case 3:
                    answer += unit[i*2] + unit[i*2] + unit[i*2]
                case 4:
                    answer += unit[i*2] + unit[i*2+1]
                case 5:
                    answer += unit[i*2+1]
                case 6:
                    answer += unit[i*2+1] + unit[i*2]
                case 7:
                    answer += unit[i*2+1] + unit[i*2] + unit[i*2]
                case 8:
                    answer += unit[i*2+1] + unit[i*2] + unit[i*2] + unit[i*2]
                case 9:
                    answer += unit[i*2] + unit[i*2+2]
            }
            num %= array[i]
        }
    }

    return answer
}

func main(){
    for i:=0; i<4000; i++{
        fmt.Printf("%15s ",intToRoman(i))
        if i%5==0{
            fmt.Println()
        }
    }
}
