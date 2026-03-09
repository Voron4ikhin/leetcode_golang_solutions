package _001_two_sum

func TwoSum(nums []int, target int) []int {
	processedMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		toFind := target - nums[i]
		index, found := processedMap[toFind]
		if found {
			return []int{index, i}
		}
		processedMap[nums[i]] = i
	}

	return nil
}
