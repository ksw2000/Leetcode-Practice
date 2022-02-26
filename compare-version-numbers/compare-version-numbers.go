// URL : https://leetcode.com/problems/compare-version-numbers/
// 0 ms 2 MB
package main

func split(src string) [500]int {
	ret := [500]int{}
	num := 0
	i := 0
	for _, c := range []rune(src) {
		if c == '.' {
			ret[i] = num
			num = 0
			i++
		} else {
			num = num*10 + int(c-'0')
		}
	}
	ret[i] = num
	return ret
}

func compareVersion(version1 string, version2 string) int {
	v1 := split(version1)
	v2 := split(version2)
	i := 0
	for ; i < len(v1) && i < len(v2); i++ {
		if v1[i] < v2[i] {
			return -1
		}
		if v1[i] > v2[i] {
			return 1
		}
	}
	for ; i < len(v1); i++ {
		if v1[i] > 0 {
			return 1
		}
	}
	for ; i < len(v2); i++ {
		if v2[i] > 0 {
			return -1
		}
	}
	return 0
}
