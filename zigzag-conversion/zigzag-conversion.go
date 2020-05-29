// https://leetcode.com/problems/zigzag-conversion/
// 4 ms	3.9 MB
package main
import "fmt"

func convert(s string, numRows int) string {
    if numRows<2 {
        return s
    }
    string_length := len(s)
    n             := 2*numRows-2
    answer        := make([]byte, string_length)
    answerIndex   := 0

    for i:=0; i<numRows; i++{
        for stringIndex:=i; stringIndex < string_length; stringIndex += n{
            answer[answerIndex] = s[stringIndex]
            answerIndex+=1
            if(i!=0 && i!=numRows-1 && stringIndex+n-i-i < string_length){
                answer[answerIndex] = s[stringIndex+n-i-i]
                answerIndex+=1
            }
        }
    }

    return (string)(answer)
}

func main(){
    fmt.Println(convert("PAYPALISHIRING",1))
    fmt.Println(convert("PAYPALISHIRING",2))
    fmt.Println(convert("PAYPALISHIRING",3))
    fmt.Println(convert("PAYPALISHIRING",4))
}
