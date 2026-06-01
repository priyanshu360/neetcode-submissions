func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		a := 0
		if i > 1 { a = dp[i - 2] }	
		dp[i] = max(dp[i - 1], nums[i - 1] + a)	
	}

	return dp[n]
}
