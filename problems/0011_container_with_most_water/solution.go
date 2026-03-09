package _001_two_sum

func MaxArea(height []int) int {
	var result int
	left := 0
	right := len(height) - 1

	for left != right {
		minHeight := 0
		if height[left] > height[right] {
			minHeight = height[right]
		} else {
			minHeight = height[left]
		}
		square := minHeight * (right - left)
		if square > result {
			result = square
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return result
}
