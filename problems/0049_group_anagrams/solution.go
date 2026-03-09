package _049_group_anagrams

import "slices"

func GroupAnagram(strs []string) [][]string {
	wordMap := make(map[string][]int)
	for i := 0; i < len(strs); i++ {
		toSortRune := []rune(strs[i])
		slices.Sort(toSortRune)
		sortedWord := string(toSortRune)
		wordMap[sortedWord] = append(wordMap[sortedWord], i)
	}

	result := make([][]string, 0, len(wordMap))
	for _, group := range wordMap {
		var r []string
		for _, i := range group {
			r = append(r, strs[i])
		}
		result = append(result, r)
	}
	return result
}
