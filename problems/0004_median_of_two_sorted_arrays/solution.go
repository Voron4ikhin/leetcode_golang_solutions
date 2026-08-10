package _004_median_of_two_sorted_arrays

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}

	len1, len2 := len(nums1), len(nums2)

	left := 0
	right := len1

	for left <= right {
		partX := (left + right) / 2
		partY := (len2+len1+1)/2 - partX

		nums1LeftMax := math.MinInt64
		if partX != 0 {
			nums1LeftMax = nums1[partX-1]
		}
		nums1RightMin := math.MaxInt64
		if partX != len1 {
			nums1RightMin = nums1[partX]
		}

		nums2LeftMax := math.MinInt64
		if partY != 0 {
			nums2LeftMax = nums2[partY-1]
		}

		nums2RightMin := math.MaxInt64
		if partY != len2 {
			nums2RightMin = nums2[partY]
		}

		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			sumLen := len1 + len2
			leftSum := partX + partY
			rightSum := sumLen - leftSum

			if rightSum > leftSum {
				return float64(min(nums2RightMin, nums1RightMin))
			}

			if leftSum > rightSum {
				return float64(max(nums1LeftMax, nums2LeftMax))
			}

			if leftSum == rightSum {
				return (float64(min(nums2RightMin, nums1RightMin)) + float64(max(nums1LeftMax, nums2LeftMax))) / 2
			}

		}

		if nums1LeftMax > nums2RightMin {
			right = partX - 1
		}

		if nums2LeftMax > nums1RightMin {
			left = partX + 1
		}
	}

	return 0.0
}
