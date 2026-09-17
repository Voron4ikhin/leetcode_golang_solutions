package _088_merge_sorted_array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{name: "Example 1", nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("Merge() = %v, want %v", tt.nums1, tt.want)
			}
		})
	}
}
