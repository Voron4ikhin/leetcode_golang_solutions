package _056_merge_intervals

import (
	"fmt"
	"sort"
)

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	merged := [][]int{intervals[0]}

	fmt.Println(merged)
	fmt.Println(intervals[1:])

	for _, current := range intervals[1:] {
		last := merged[len(merged)-1]
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}
