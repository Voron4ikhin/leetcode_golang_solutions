package _228_summary_ranges

import (
	"strconv"
)

func SummaryRanges(nums []int) []string {
	res := []string{}

	if len(nums) == 0 {
		return res
	}

	start, end := 0, 0
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(nums[start]))
		return res
	}
	//0, 1, 2, 4, 5,
	for i := 1; i < len(nums); i++ {
		if nums[end]+1 == nums[i] {
			end++
			continue
		}
		if start == end {
			res = append(res, strconv.Itoa(nums[start]))
		} else {
			res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
		}
		start, end = i, i
	}

	if start == end {
		res = append(res, strconv.Itoa(nums[start]))
		return res
	}

	res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))

	return res
}
