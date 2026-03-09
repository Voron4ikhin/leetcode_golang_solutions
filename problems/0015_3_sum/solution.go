package _049_group_anagrams

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for index := 0; index < len(nums); index++ {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		left := index + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if sum == 0 {
				preResult := []int{nums[index], nums[left], nums[right]}
				result = append(result, preResult)
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
