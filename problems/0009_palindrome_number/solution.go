package _009_palindrome_number

import "strconv"

func IsPalindrome(x int) bool {
	strInt := strconv.Itoa(x)
	left, right := 0, len(strInt)-1
	for left < right {
		if strInt[left] != strInt[right] {
			return false
		}
		left++
		right--
	}

	return true
}
