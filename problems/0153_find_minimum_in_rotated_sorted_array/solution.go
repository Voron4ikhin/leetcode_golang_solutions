package _153_find_minimum_in_rotated_sorted_array

func FindMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2

		if nums[mid] > nums[high] {
			// идем в сторону high
			low = mid + 1
		} else {
			// идем в строну low
			high = mid
		}
	}

	return nums[low]
}
