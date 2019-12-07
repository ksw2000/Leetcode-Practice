// URL : https://leetcode.com/problems/integer-to-roman/
//0 ms	3.3 MB
package main
import "fmt"

func intToRoman(num int) string {
    array  := []int{1,10,100,1000}
    unit   := []byte{'I','V','X','L','C','D','M'}
    answer := make([]byte,17)
    index  := 0
    for i := len(array)-1; i>=0; i--{
        for num/array[i]>0 {
            switch num/array[i] {
                case 1:
                    answer[index]=unit[i*2]
                    index += 1
                case 2:
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                case 3:
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                case 4:
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2+1]
                    index += 1
                case 5:
                    answer[index]=unit[i*2+1]
                    index += 1
                case 6:
                    answer[index]=unit[i*2+1]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                case 7:
                    answer[index]=unit[i*2+1]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                case 8:
                    answer[index]=unit[i*2+1]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2]
                    index += 1
                case 9:
                    answer[index]=unit[i*2]
                    index += 1
                    answer[index]=unit[i*2+2]
                    index += 1
            }
            num %= array[i]
        }
    }

    return string(answer[0:index])
}

func main(){
    for i:=0; i<4000; i++{
        fmt.Printf("%15s ",intToRoman(i))
        if i%5==0{
            fmt.Println()
        }
    }
}
