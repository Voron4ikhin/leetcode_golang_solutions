package _003_longest_substring_without_repeating_characters

func LengthOfLongestSubstring(s string) int {
	var result int
	var left int
	//Слайс указывающий когда символ встречался последний раз в строке
	lastIndexSlice := make([]int, 128)

	runeString := []rune(s)
	for right := 0; right < len(runeString); right++ {
		if lastIndexSlice[runeString[right]] > left {
			left = lastIndexSlice[runeString[right]]
		}
		if right-left+1 > result {
			result = right - left + 1
		}
		lastIndexSlice[runeString[right]] = right + 1
	}

	return result
}
