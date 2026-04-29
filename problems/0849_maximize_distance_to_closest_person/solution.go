package _849_maximize_distance_to_closest_person

func MaxDistToClosest(seats []int) int {
	n := len(seats)
	maxGap := 0
	start := -1

	for i := range seats {
		if seats[i] == 0 {
			continue
		}
		if start == -1 {
			maxGap = i
		} else {
			maxGap = max(maxGap, (i-start)/2)
		}
		start = i
	}

	return max(maxGap, n-1-start)
}
