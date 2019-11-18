//https://leetcode.com/problems/longest-common-prefix/
//0 ms	2.4 MB
package main
import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs)==0{
        return ""
    }
    min_len := len(strs[0])
    len_of_strs := len(strs)

    for i:=1; i<len_of_strs; i++{
        if len(strs[i]) < min_len{
            min_len=len(strs[i])
        }
    }

    breakFlag:=0
    length:=0
    for i:=0; i<min_len; i++ {
        for j:=0; j<len_of_strs; j++{
            if strs[j][i]!=strs[0][i] {
                breakFlag=1
            }
        }
        if breakFlag == 0{
            length++
        }else{
            break
        }
    }
    return strs[0][0:length]
}

func main(){
    var strs = []string{"flower","flow","flight"}
    fmt.Println(longestCommonPrefix(strs))
    var strs2 = []string{}
    fmt.Println(longestCommonPrefix(strs2))
    //str := "fuckyou"
    //fmt.Println(str[2:4])
}
