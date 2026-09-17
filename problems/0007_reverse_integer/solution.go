package _007_reverse_integer

import "math"

func Reverse(x int) int {
	var res int32
	var maxInt32 int32
	maxInt32 = math.MaxInt32
	for x != 0 {
		ost := x % 10
		if maxInt32/10 < res {
			return 0
		}
		res = res*10 + int32(ost)
		x = x / 10
	}

	return int(res)
}
