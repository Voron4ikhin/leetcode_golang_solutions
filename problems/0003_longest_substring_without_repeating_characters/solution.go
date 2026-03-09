package _003_longest_substring_without_repeating_characters

func LengthOfLongestSubstring(s string) int {
	var result int
	var left int

	letterMap := make(map[rune]struct{})
	runeString := []rune(s)
	for right := 0; right < len(runeString); right++ {
		if _, ok := letterMap[runeString[right]]; ok {
			for runeString[left] != runeString[right] {
				delete(letterMap, runeString[left])
				left++
			}
			delete(letterMap, runeString[left])
			left++
		}

		letterMap[runeString[right]] = struct{}{}
		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
