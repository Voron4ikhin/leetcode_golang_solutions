package _070_climbing_stairs

func ClimbStairs(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func badClimbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	left := ClimbStairs(n - 1)
	right := ClimbStairs(n - 2)

	return left + right
}
