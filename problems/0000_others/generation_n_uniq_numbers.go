package _000_others

import "math/rand"

// Требуется реализовать функцию uniqRandn, которая генерирует слайс длины n уникальных, рандомных чисел.
func UniqRandn(n int) []int {
	result := make([]int, 0, n)
	resMap := make(map[int]struct{}, n)

	for len(result) < n {
		val := rand.Intn(101)
		if _, ok := resMap[val]; ok {
			continue
		}
		result = append(result, val)
		resMap[val] = struct{}{}
	}
	return result
}
