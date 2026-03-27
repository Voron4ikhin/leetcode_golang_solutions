package _242_valid_anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}

func IsAnagramBad(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	checkMap := make(map[rune]int)
	for _, v := range s {
		checkMap[v]++
	}
	for _, v := range t {
		if _, ok := checkMap[v]; !ok {
			return false
		}
		checkMap[v]--
		if checkMap[v] < 0 {
			return false
		}
	}

	return true
}
