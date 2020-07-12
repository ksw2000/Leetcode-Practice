// URL : https://leetcode.com/problems/longest-substring-without-repeating-characters/
// 8 ms 3.1 MB
// Dynamic Programming
package main
import "fmt"
func lengthOfLongestSubstring(s string) int {
    start:=0
    maxLen:=-1
    seen:=make(map[byte]int)
    for i , value := range s{   // see i as end of slice
        if(seen[byte(value)] > start){
            //If this slice cotain repeat character, then change start
            start=seen[byte(value)]
        }
        seen[byte(value)]=i+1   //Storage position+1 into seen[char]
        fmt.Println(start,i)    //print the slice start and end
        if(i-start>maxLen){
            maxLen=i-start
        }
    }
    return maxLen+1
}

func main(){
    fmt.Printf("%s %d\n","",lengthOfLongestSubstring(""))
    fmt.Printf("%s %d\n","abcabcbb",lengthOfLongestSubstring("abcabcbb"))
    fmt.Printf("%s %d\n","bbbbb",lengthOfLongestSubstring("bbbbb"))
    fmt.Printf("%s %d\n","pwwkew",lengthOfLongestSubstring("pwwkew"))
}
