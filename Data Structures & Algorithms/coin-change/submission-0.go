func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	const INF = 1_000_000_009
	for i := range dp {
		dp[i] = INF
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i { continue }
			dp[i] = min(dp[i], dp[i - coin] + 1)
		}
	}

	if dp[amount] == INF {return -1} else {return dp[amount]}
}
