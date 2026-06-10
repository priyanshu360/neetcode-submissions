func minDistance(word1 string, word2 string) int {
	const INF = 1_000_000_000

	n, m := len(word1), len(word2)
	dp := make([][]int, n + 1)

	for i := range n + 1 {
		dp[i] = make([]int, m + 1)
		for j := range m + 1 {
			dp[i][j] = INF
			if i == 0 || j == 0 {
				dp[i][j] = max(i, j)
			}
		}
	}

	//abc
	//axy

	for i, v := range word1 {
		for j, w := range word2 {
			dp[i + 1][j + 1] = min(1 + dp[i][j + 1], 1 + dp[i + 1][j], 1 + dp[i][j])
			if v == w {
				dp[i + 1][j + 1] = min(dp[i + 1][j + 1], dp[i][j])
			}
		} 
	}


	return dp[n][m]
}
