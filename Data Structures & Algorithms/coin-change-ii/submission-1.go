func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for amt := coin; amt <= amount; amt++ {
			dp[amt] += dp[amt-coin]
		}
	}

	return dp[amount]
}
