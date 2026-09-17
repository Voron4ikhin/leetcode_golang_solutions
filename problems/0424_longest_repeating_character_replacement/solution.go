package _424_longest_repeating_character_replacement

func CharacterReplacement(s string, k int) int {
	counts := [26]int{}
	l := 0
	maxFreq := 0
	res := 0

	for r := 0; r < len(s); r++ {
		idx := s[r] - 'A'
		counts[idx]++

		maxFreq = max(maxFreq, counts[idx])

		for (r-l+1)-maxFreq > k {
			counts[s[l]-'A']--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
