// URL : https://leetcode.com/problems/longest-palindromic-substring/submissions/
// 8 ms	2.8 MB
package main
import "fmt"

func longestPalindrome(s string) string {
    //Mark #
    old_s := []rune(s)
    new_s := make([]rune, 2*len(s)+1)
    new_s[0] = '#'
    for i:=0; i<len(s); i++{
        new_s[i*2+1] = old_s[i]
        new_s[i*2+2] = '#'
    }

    //Count palindromic radius
    max, max_index := 0, 0   //max: max of radius, max_index: the index_of_max
    for i := range(new_s){
        count := 0
        for j:=1; i-j>=0 && i+j<len(new_s); j++{
            if new_s[i+j] == new_s[i-j]{
                count++
            }else{
                break
            }
        }

        if count>max{
            max, max_index = count, i
        }
    }

    //Get answer according to max and max_index
    if(max_index % 2 == 1){  // center on alphabet (the length of substring is odd)
        max = (max-1)>>1
        max_index = (max_index-1)>>1
        return s[max_index - max : max_index + max +1]
    }else{  //center on # (the length of substring is even)
        max = max>>1
        max_index = max_index>>1
        return s[max_index - max : max_index + max]
    }
    return ""
}

func main(){
    fmt.Println(longestPalindrome("a"))
    fmt.Println(longestPalindrome("cc"))
    fmt.Println(longestPalindrome("abb"))
    fmt.Println(longestPalindrome("aabaa"))
    fmt.Println(longestPalindrome("abbccbba"))
}
