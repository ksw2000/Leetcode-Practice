// https://leetcode.com/problems/simplify-path/
// 0 ms 2.8 MB
package main

import "strings"

func simplifyPath(path string) string {
	pathList := strings.Split(path, "/")
	j := 0
	for i, path := range pathList {
		if path == ".." {
			if j > 0 {
				j--
			}
		} else if path != "" && path != "." {
			if i != j {
				pathList[j] = pathList[i]
			}
			j++
		}
	}
	return "/" + strings.Join(pathList[:j], "/")
}
