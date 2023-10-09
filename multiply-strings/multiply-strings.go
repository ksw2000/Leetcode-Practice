// https://leetcode.com/problems/multiply-strings/
package main

func multiply(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)

	result := make([]byte, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			result[i+j+1] = result[i+j+1] + (num1[i]-'0')*(num2[j]-'0')
			if result[i+j+1] >= 10 {
				v := result[i+j+1]
				result[i+j+1] = v % 10
				v = v / 10
				for k := i + j; v > 0; k-- {
					v = result[k] + v
					result[k] = v % 10
					v = v / 10
				}
			}
		}
	}

	begin := 0
	for ; begin < len(result)-1 && result[begin] == 0; begin++ {
	}
	result = result[begin:]

	for i := range result {
		result[i] += '0'
	}

	return string(result)
}
