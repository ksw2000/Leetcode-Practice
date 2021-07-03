package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	input := countAndSay(n - 1)
	node := -1
	var ret strings.Builder
	for i := range input {
		if i == len(input)-1 || input[i] != input[i+1] {
			ret.WriteString(strconv.Itoa(i - node))
			ret.WriteByte(input[i])
			node = i
		}
	}
	return ret.String()
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))
}
