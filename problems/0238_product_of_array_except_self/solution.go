package _238_product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	zeroIndex := -1
	multiAll := 1

	for k, v := range nums {
		if v == 0 {
			if zeroIndex != -1 {
				return make([]int, len(nums))
			}
			zeroIndex = k
			continue
		}
		multiAll *= v
	}

	for i, num := range nums {
		if zeroIndex == -1 {
			result[i] = multiAll / num
			continue
		}
		if i != zeroIndex {
			result[i] = 0
			continue
		}
		result[i] = multiAll
	}

	return result
}
