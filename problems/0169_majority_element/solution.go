package _169_majority_element

func MajorityElement(nums []int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}

	for k, v := range numsMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
