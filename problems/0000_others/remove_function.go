package _000_others

import "fmt"

// Дан слайс целых чисел. Напишите функцию remove которая удаляет из слайса все нули

// 1, 3, 0, 7, 9, 0, 12, 0, 0, 13, 15
func Remove(nums []int) []int {
	var i, j int // j показывает индекс нуля
	for i < len(nums) {
		fmt.Println("i, j before")
		fmt.Println(i, j)
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
		fmt.Println("i, j after")
		fmt.Println(i, j)
	}

	return nums[:j]
}
