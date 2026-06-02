func findTargetSumWays(nums []int, target int) int {
	const offset = 1000

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, 2001)
	}

	dp[0][offset] = 1

	for i, num := range nums {
		for sum := 0; sum <= 2000; sum++ {
			if dp[i][sum] == 0 {
				continue
			}

			dp[i+1][sum+num] += dp[i][sum]
			dp[i+1][sum-num] += dp[i][sum]
		}
	}

	if target < -1000 || target > 1000 {
		return 0
	}

	return dp[len(nums)][target+offset]
}
