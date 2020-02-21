// URL : https://leetcode.com/problems/generate-parentheses/
// 0 ms	2.8 MB
package main
import "fmt"

//----------------------------------- Answer Begin -----------------------------------//
func generate(left int, right int, n *int, result string, answer *[]string) {
    if(left + right == (*n)<<1){
        *answer = append(*answer, result)
    }else{
        if(left > right){
            if(left < *n){
                generate(left+1, right, n, result+"(", answer)
            }
            generate(left, right+1, n, result+")", answer)
        }else if(left == right){
            generate(left+1, right, n, result+"(", answer)
        }
    }
}

func generateParenthesis(n int) []string {
    answer := []string{}
    if n>0{
        generate(1, 0, &n, "(", &answer)
    }
    return answer
}
//----------------------------------- Answer END -----------------------------------//

func main(){
    fmt.Println(generateParenthesis(0))
    fmt.Println(generateParenthesis(1))
    fmt.Println(generateParenthesis(2))
    fmt.Println(generateParenthesis(3))
}
