package _136_single_number

func SingleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	return res
}

func BadSingleNumber(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		_, ok := numsMap[nums[i]]
		if !ok {
			numsMap[nums[i]] = struct{}{}
			continue
		}
		delete(numsMap, nums[i])
	}

	for index, _ := range numsMap {
		return index
	}

	return -1
}
