// substring-with-concatenation-of-all-words
// Time: O(len(s) * len(words))
// Space: O(3*len(words) + len(s))
package main

func findSubstring(s string, words []string) []int {
	invertedIndex := map[string]int{}
	groundTruth := make([]int, len(words)+1)
	groundTruthSum := len(words)

	for i := 0; i < len(words); i++ {
		if j, ok := invertedIndex[words[i]]; ok {
			groundTruth[j+1]++
		} else {
			invertedIndex[words[i]] = len(invertedIndex)
			groundTruth[i+1]++
		}
	}

	groundTruth = groundTruth[:len(invertedIndex)+1]

	mark := make([]int, len(s))
	substringLen := len(words[0])
	for i := 0; i <= len(s)-substringLen; i++ {
		if j, ok := invertedIndex[s[i:i+substringLen]]; ok {
			mark[i] = j + 1
		}
	}

	// fmt.Println(words)
	// fmt.Println(groundTruth)
	// fmt.Println(mark)

	visited := make([]int, len(groundTruth))
	ret := []int{}

forMark:
	for i := 0; i < len(mark); i++ {
		// initialize visited
		for j := range visited {
			visited[j] = groundTruth[j]
		}

		k := 0
		for j := i; j < len(mark) && k < groundTruthSum; j += substringLen {
			visited[mark[j]]--
			k++
			if visited[mark[j]] < 0 {
				continue forMark
			}
		}

		if k == groundTruthSum {
			ret = append(ret, i)
		}
	}

	return ret
}
