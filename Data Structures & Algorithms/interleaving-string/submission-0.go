func isInterleave(s1 string, s2 string, s3 string) bool {
	x, y, z := len(s1), len(s2), len(s3)

	if x + y != z {
		return false
	}

	dp := make([][]bool, x+1)
	for i := range dp {
		dp[i] = make([]bool, y+1)
	}

	dp[0][0] = true

	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			if i > 0 {
				if s1[i - 1] == s3[i+j-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]		
				}
			}
			if j > 0 {
				if s2[j - 1] == s3[i+j-1] {
					dp[i][j] = dp[i][j] || dp[i][j-1]
				}
			}
		}
	}

	return dp[x][y]
}
