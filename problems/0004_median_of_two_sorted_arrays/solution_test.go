package _004_median_of_two_sorted_arrays

import (
	"reflect"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 3, 5, 7, 9},
			nums2: []int{2, 4, 6, 8, 10},
			want:  5.5,
		},
		{
			name:  "Example 2",
			nums1: []int{1, 3, 5, 7, 9},
			nums2: []int{2, 4, 6, 8},
			want:  5,
		},
		{
			name:  "Example 3",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{6, 7, 8, 9, 10},
			want:  5.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianSortedArrays(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMedianSortedArrays result is = %v, want %v", got, tt.want)
			}
		})
	}
}
