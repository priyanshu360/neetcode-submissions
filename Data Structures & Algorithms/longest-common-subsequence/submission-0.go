func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	dp := make([][]int, n + 1)
	for i := range n+1 {	
		dp[i] = make([]int, m + 1)

		for j := range m + 1 {
			if i == 0 ||  j == 0 {
				continue
			} 

			match := 0
			if text1[i - 1] == text2[j - 1] { match = 1 }

			dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

			if match == 1 {
				dp[i][j] = max(dp[i][j], dp[i - 1][j - 1] + 1)
			}
		}
	}

	return dp[n][m]	
}
