// URL : https://leetcode.com/problems/longest-substring-without-repeating-characters/
// 92 ms 2.6 MB
// Brute Force
package main
import "fmt"
func lengthOfLongestSubstring(s string) int {
    if s==""{
        return 0
    }
    maxLen:=1
    var breakFlag bool
    for i:=0; i<len(s)-maxLen; i++{
        for j:=1; i+j < len(s); j++{
            breakFlag=false
            for _ , value := range s[i:i+j]{
                if byte(value) == byte(s[i+j]){
                    breakFlag=true
                    break
                }
            }
            if breakFlag==true{
                break
            }else{
                if j+1>maxLen{
                    maxLen = j+1
                    //fmt.Println(s[i:i+j+1])
                }
            }
        }
    }
    return maxLen
}

func main(){
    fmt.Printf("%s %d\n","",lengthOfLongestSubstring(""))
    fmt.Printf("%s %d\n","abcabcbb",lengthOfLongestSubstring("abcabcbb"))
    fmt.Printf("%s %d\n","bbbbb",lengthOfLongestSubstring("bbbbb"))
    fmt.Printf("%s %d\n","pwwkew",lengthOfLongestSubstring("pwwkew"))
}
