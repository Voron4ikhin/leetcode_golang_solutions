package _053_maximum_subarray

func MaxSubArray(nums []int) int {
	current := nums[0]
	maximum := nums[0]

	for i := 1; i < len(nums); i++ {
		current = max(nums[i], current+nums[i])
		maximum = max(maximum, current)
	}

	return maximum
}
