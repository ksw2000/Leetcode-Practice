// URL : https://leetcode.com/problems/valid-parentheses
//0 ms	2 MB
package main
import(
    "fmt"
    "os"
)

type stack struct{
    token int
    next *stack
}

func getToken(c byte) int{
    switch c {
        case '(':
            return 1
        case ')':
            return -1
        case '[':
            return 2
        case ']':
            return -2
        case '{':
            return 3
        case '}':
            return -3
    }
    return 0
}

func push(s **stack, val int){
    var newStack *stack
    newStack=new(stack)
    (*newStack).next=(*s)
    (*newStack).token=val
    (*s)=newStack
}

func pop(s **stack) int{
    if (*s) == nil{
        fmt.Fprintf(os.Stderr, "error: %v\n", (*s))
        os.Exit(1)
    }
    token := (**s).token
    (*s) = (**s).next
    return token
}

func isValid(s string) bool {
    var getTokenInt int
    var rec *stack
    len_of_s:=len(s)
    push(&rec,0)
    for i:=0; i<len_of_s; i++{
        getTokenInt=getToken(s[i])
        if ((*rec).token + getTokenInt) == 0 {
            pop(&rec)
        }else if ((*rec).token * getTokenInt) < 0{
            return false
        }else{
            push(&rec, getTokenInt)
        }
    }
    if pop(&rec) == 0{
        return true
    }
    return false
}

func main(){
    s:="()"
    fmt.Printf("%s \t\t %t\n",s,isValid(s))
    s="()[]{}"
    fmt.Printf("%s \t\t %t\n",s,isValid(s))
    s="(]"
    fmt.Printf("%s \t\t %t\n",s,isValid(s))
    s="{[]}"
    fmt.Printf("%s \t\t %t\n",s,isValid(s))
}
