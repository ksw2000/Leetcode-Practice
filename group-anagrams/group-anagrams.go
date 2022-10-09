// https://leetcode.com/problems/group-anagrams/
// 24 ms 7.8 MB

package main

func generateKey(s string) (key [26]int) {
	for _, r := range s {
		key[r-'a']++
	}
	return key
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	set := map[[26]int]int{}
	for _, str := range strs {
		key := generateKey(str)
		if index, existed := set[key]; !existed {
			set[key] = len(result)
			result = append(result, []string{str})
		} else {
			result[index] = append(result[index], str)
		}
	}
	return result
}
