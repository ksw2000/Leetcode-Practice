// 0 ms	2.4 MB
// URL : https://leetcode.com/problems/multiply-strings/
package main
import "fmt"

func reverse(str []rune) []rune{
    for i:=0; i<len(str)/2; i++{
        str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
    }
    return str
}

func multiply(num1 string, num2 string) string {
    var ret string
    if(len(num1)>0 && len(num2)>0){
        tmpInt := make([]int,len(num1)+len(num2))
        runeNum1,runeNum2 := reverse([]rune(num1)),reverse([]rune(num2))
        for i:=0; i<len(runeNum1); i++{
            for j:=0; j<len(runeNum2); j++{
                tmpInt[i+j] += int(runeNum1[i]-'0') * int(runeNum2[j]-'0')
            }
        }

        for index, val := range(tmpInt){
            if val>=10{
                tmpInt[index] = val % 10
                val /= 10
                for i:=index+1 ; val>0 ; i++{
                    tmpInt[i] += val % 10
                    val/=10
                }
            }
        }

        end := 0
        for index, val := range(tmpInt){
            if val>0{
                end = index
            }
        }

        tmpInt = tmpInt[0:end+1]
        tmpRune := make([]rune, end+1)
        for i:=0; i<=end; i++{
            tmpRune[i] = rune(tmpInt[i]+int('0'))
        }
        ret = string(reverse(tmpRune))
    }

    return ret
}

func main(){
    num1 := "5"
    num2 := "12"
    fmt.Println(multiply(num1, num2))
}
