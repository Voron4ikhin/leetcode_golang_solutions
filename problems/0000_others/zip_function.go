package _000_others

// Требуется реализовать функцию zip, которая соединяет элементы двух слайсов в слайс пар
func Zip(s ...[]int) [][]int {
	if len(s) == 0 {
		return [][]int{}
	}
	minLength := len(s[0])
	for i := 1; i < len(s); i++ {
		if len(s[i]) < minLength {
			minLength = len(s[i])
		}
	}

	result := make([][]int, 0, minLength)

	for i := 0; i < minLength; i++ {
		preResult := make([]int, 0, len(s))
		for j := 0; j < len(s); j++ {
			preResult = append(preResult, s[j][i])
		}
		result = append(result, preResult)
	}

	return result
}
