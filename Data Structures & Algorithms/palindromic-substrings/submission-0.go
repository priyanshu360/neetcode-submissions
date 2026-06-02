func countSubstrings(s string) int {
	var ans int
	n := len(s)
	dp := make([][]int, n) 

	for i := range n {
		dp[i] = make([]int, n)

		dp[i][i] = 1
		ans++
	}


	for l := 2; l <= n; l++ {
		for i := 0; i <= n - l; i++ {
			if s[i] == s[i + l - 1] {
				if l == 2 || l == 3 {
					dp[i][i + l - 1] = l
					ans++
				} else {
					if dp[i+1][i + l - 2] > 1 {
						dp[i][i + l - 1] = 2 + dp[i+1][i + l - 2]
						ans++
					}
				}
			}	
		}
	}
	
	return ans
}
