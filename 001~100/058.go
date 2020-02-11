// URL : https://leetcode.com/problems/length-of-last-word/
// 0 ms	2.2 MB
package main
import "fmt"

func lengthOfLastWord(s string) int {
    count := 0
    count_space := 0
    for j:=len(s)-1; j>=0; j--{
        if s[j]==' '{
            count_space++;
        }else{
            break
        }
    }

    for i:=0; i<len(s)-count_space; i++{
        if s[i]==' '{
            count=0
        }else{
            count++
        }
    }
    return count
}

func main(){
    fmt.Println(lengthOfLastWord("Hello World"))
    fmt.Println(lengthOfLastWord("Hello"))
    fmt.Println(lengthOfLastWord("Hello "))
    fmt.Println(lengthOfLastWord(""))
}
