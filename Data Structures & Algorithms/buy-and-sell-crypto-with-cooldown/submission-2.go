func maxProfit(prices []int) int {
//	dp[i] = (dp[j - 2] + profit with j) 	

	n := len(prices)
	dp := make([]int, n)

	for i := range prices {
		if i != 0 {dp[i] = dp[i - 1]}
		for j := range i {
			val := 0
			if j > 2 {
				val = dp[j - 2]
			}
			dp[i] = max(dp[i], val + prices[i] - prices[j])
		}
	}

	return dp[n - 1]
}
