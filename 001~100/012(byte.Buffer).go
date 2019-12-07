// URL : https://leetcode.com/problems/integer-to-roman/
//8 ms	3.7 MB
package main
import (
    "fmt"
    "bytes"
)

//改用 bytes.Buffer 加速似乎沒有比較快
func intToRoman(num int) string {
    array  := []int   {1,10,100,1000}
    unit   := []byte{'I','V','X','L','C','D','M'}
    var answer bytes.Buffer
    for i := len(array)-1; i>=0; i--{
        for num/array[i]>0 {
            switch num/array[i] {
                case 1:
                    answer.WriteByte(unit[i*2])
                case 2:
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2])
                case 3:
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2])
                case 4:
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2+1])
                case 5:
                    answer.WriteByte(unit[i*2+1])
                case 6:
                    answer.WriteByte(unit[i*2+1])
                    answer.WriteByte(unit[i*2])
                case 7:
                    answer.WriteByte(unit[i*2+1])
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2])
                case 8:
                    answer.WriteByte(unit[i*2+1])
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2])
                case 9:
                    answer.WriteByte(unit[i*2])
                    answer.WriteByte(unit[i*2+2])
            }
            num %= array[i]
        }
    }

    return answer.String()
}

func main(){
    for i:=0; i<4000; i++{
        fmt.Printf("%15s ",intToRoman(i))
        if i%5==0{
            fmt.Println()
        }
    }
}
