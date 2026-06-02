func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n + 1) 


	isValid := func(s string) bool {
		switch s[0] {
			case '1' :
				return true
			case '2' :
				return s[1] < '7'
			default  :
				return false
		}
	}

	dp[0] = 1
	if s[0] != '0' { dp[1] = 1 }
	for i := 2; i <= n; i++ {
		if s[i - 1] != '0' {
			if dp[i - 1] != 0 { dp[i] = dp[i - 1] }	
		}

		if dp[i - 1] != 0 && isValid(s[i-2:i]){
			dp[i] += dp[i - 2]
		}	
		
	}

	return dp[n]
}
