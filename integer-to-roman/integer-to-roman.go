// https://leetcode.com/problems/integer-to-roman/
// 7 ms	3.3 MB

package main

var mapping = [...]struct {
	arabic int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	ret := ""
	for i := 0; num > 0; {
		if num >= mapping[i].arabic {
			num -= mapping[i].arabic
			ret += mapping[i].roman
		} else {
			i++
		}
	}
	return ret
}
