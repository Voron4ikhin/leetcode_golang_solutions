package _347_top_k_frequent_elements

import "fmt"

func TopKFrequent(nums []int, k int) []int {
	n := len(nums)
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	fmt.Println(freqMap)

	buckets := make([][]int, n+1)

	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}
	fmt.Println(buckets)

	res := []int{}

	for i := n; i >= 0; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)

			if len(res) == k {
				fmt.Println(res)
				return res
			}
		}
	}

	return res
}
