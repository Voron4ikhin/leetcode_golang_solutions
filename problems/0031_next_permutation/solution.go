package _031_next_permutation

import (
	"slices"
)

func NextPermutation(nums []int) []int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}

	if index != 0 {
		slices.Sort(nums[index:])
		Swap(nums, index)
	} else {
		slices.Sort(nums)
	}

	return nums
}

func Swap(nums []int, i int) {
	indexToSwap := i - 1
	for nums[i] < nums[indexToSwap] {
		i++
	}
	value := nums[i]
	nums[i] = nums[indexToSwap]
	nums[indexToSwap] = value
}
