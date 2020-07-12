// URL : https://leetcode.com/problems/string-to-integer-atoi/
// 0 ms 2.3 MB
package main
import "fmt"

func myAtoi(str string) int {
    len_of_str := len(str)
    if len_of_str==0{
        return 0
    }

    i:=0
    for ; i<len_of_str-1; i++ {
        if str[i]!=' '{
            break;
        }
    }

    negation := false
    if str[i]=='-'{
        negation = true
        i++
    }else if str[i]=='+'{
        negation = false
        i++
    }else if str[i]<'0' || str[i]>'9' {
        return 0
    }

    var tmp int64
    for j:=i; j<len_of_str; j++ {
        if str[j]>='0' && str[j]<='9'{
            tmp = tmp*10 + int64(str[j]) - '0'
        }else{
            break
        }
        //Overflow
        if negation && tmp >= 2147483648{
            return -2147483648
        }else if !negation && tmp>= 2147483647{
            return 2147483647
        }
    }

    if negation{
        return int(tmp*(-1))
    }else{
        return int(tmp)
    }
}

func main(){
    fmt.Printf("myAtoi(\"%s\") = %d\n", "   -42", myAtoi("   -42"))
    fmt.Printf("myAtoi(\"%s\") = %d\n", "4193 with words", myAtoi("4193 with words"))
    fmt.Printf("myAtoi(\"%s\") = %d\n", "words and 987", myAtoi("words and 987"))
    fmt.Printf("myAtoi(\"%s\") = %d\n", "-91283472332", myAtoi("-91283472332"))
    fmt.Printf("myAtoi(\"%s\") = %d\n", "3.1415", myAtoi("3.1415"))
    fmt.Printf("myAtoi(\"%s\") = %d\n", "", myAtoi(""))
    fmt.Printf("myAtoi(\"%s\") = %d\n", " ", myAtoi(" "))
    fmt.Printf("myAtoi(\"%s\") = %d\n", "+1", myAtoi("+1"))
}
