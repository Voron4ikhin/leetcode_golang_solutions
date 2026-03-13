package _075_sort_colors

func SortColors(nums []int) []int {
	countRed := 0
	countWhite := 0
	countBlue := 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			countRed++
			continue
		case 1:
			countWhite++
			continue
		case 2:
			countBlue++
			continue
		}
	}

	for i := 0; i < len(nums); i++ {
		if countRed > 0 {
			nums[i] = 0
			countRed--
			continue
		}
		if countWhite > 0 {
			nums[i] = 1
			countWhite--
			continue
		}
		if countBlue > 0 {
			nums[i] = 2
			countBlue--
			continue
		}
	}

	return nums
}
