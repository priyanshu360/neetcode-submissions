func isMatch(s string, p string) bool {
	n, m := len(s), len(p)

	dp := make([][]bool, n + 1)

	for i := range dp {
		dp[i] = make([]bool, m + 1)
	}

	dp[0][0] = true

	for j := 1; j <= m; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j-1] // zero occurrences
				if p[j-1] == '.' || p[j-1] == s[i] {
					dp[i+1][j+1] = dp[i+1][j+1] || dp[i][j+1]
				}
			}
		}
	}

	return dp[n][m]
}
