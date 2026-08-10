package _005_longest_palindromic_substring

func LongestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		opt1, opt2 := getPal(s, i, i), getPal(s, i, i+1)
		if len(opt1) > len(res) {
			res = opt1
		}
		if len(opt2) > len(res) {
			res = opt2
		}
	}

	return res
}

func getPal(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}
