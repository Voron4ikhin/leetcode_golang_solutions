package _217_contains_duplicate

func ContainsDuplicate(nums []int) bool {
	duplicateMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := duplicateMap[num]; ok {
			return true
		}
		duplicateMap[num] = struct{}{}
	}

	return false
}
