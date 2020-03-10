// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// 0 ms	2 MB
package main
import "fmt"

func concatenation(length int, digits []rune, m map[rune][]rune, answer []string) []string{
    if length > 0{
        newanswer := []string{}
        for _, val := range m[digits[0]]{
            for _,val2 := range answer{
                val2 += string(val)
                newanswer = append(newanswer, val2)
            }
        }
        return concatenation(length-1, digits[1:length], m, newanswer)
    }
    return answer
}

func letterCombinations(digits string) []string {
    m := make(map[rune][]rune)
    m['2'] = []rune{'a','b','c'}
    m['3'] = []rune{'d','e','f'}
    m['4'] = []rune{'g','h','i'}
    m['5'] = []rune{'j','k','l'}
    m['6'] = []rune{'m','n','o'}
    m['7'] = []rune{'p','q','r','s'}
    m['8'] = []rune{'t','u','v'}
    m['9'] = []rune{'w','x','y','z'}

    if len(digits) == 0{
        return []string{}
    }
    return concatenation(len(digits), []rune(digits), m, []string{""})
}

func main(){
    fmt.Println(letterCombinations(""))
    fmt.Println(letterCombinations("23"))
}
