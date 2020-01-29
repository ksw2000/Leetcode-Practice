// URL : https://leetcode.com/problems/add-binary/
// 0 ms	2.6 MB
package main
import "fmt"

//----------------------------------- Answer Begin -----------------------------------//
func cti(a byte) int{
    return int(a)-48
}

func addBinary(a string, b string) string {
    lastA:=len(a)-1
    lastB:=len(b)-1
    var answer string

    if lastA < lastB{
        return addBinary(b,a)
    }
    if lastA==-1 && lastB==-1{
        return "0"
    }else if lastA==-1{
        return b
    }else if lastB==-1{
        return a
    }

    carry:=0
    for i:=0; lastB-i>=0; i++ {
        answer = string(byte(cti(a[lastA-i]) ^ cti(b[lastB-i]) ^ carry)+48) + answer
        carry = (cti(a[lastA-i]) & cti(b[lastB-i])) | (cti(a[lastA-i]) & carry) | (carry & cti(b[lastB-i]))

    }

    for i:=0; lastA-lastB-1-i >= 0; i++ {
        answer = string(byte(cti(a[lastA-lastB-1-i]) ^ 0 ^ carry)+48) + answer
        carry = (cti(a[lastA-lastB-1-i]) & carry)

    }

    if carry==1{
        answer = "1" + answer
    }

    return answer
}
//----------------------------------- Answer END -----------------------------------//

func main(){
    fmt.Println(addBinary("",""))
}
