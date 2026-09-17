package _001_two_sum

func MaxArea(height []int) int {
	var result int
	left, right := 0, len(height)-1
	for left < right {
		areaPre := area(left, right, height)
		if areaPre > result {
			result = areaPre
		}
		if height[left] < height[right] {
			left++
			continue
		}
		if height[left] >= height[right] {
			right--
			continue
		}
	}

	return result
}

func area(l, r int, height []int) int {
	return (r - l) * min(height[l], height[r])
}
