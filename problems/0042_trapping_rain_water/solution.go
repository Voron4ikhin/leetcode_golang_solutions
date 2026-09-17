package _042_trapping_rain_water

func Trap(height []int) int {
	l := 0
	r := len(height) - 1
	maxL := height[l]
	maxR := height[r]
	water := 0

	for l < r {
		if maxR > maxL {
			l++
			maxL = max(maxL, height[l])
			water += maxL - height[l]
		} else {
			r--
			maxR = max(maxR, height[r])
			water += maxR - height[r]
		}
	}
	return water
}
