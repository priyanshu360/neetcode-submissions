func wordBreak(s string, wordDict []string) bool {
	slen := len(s)
	dp := make([]bool, slen + 1)

	dp[0] = true

	for i := range s {
		for _, word := range wordDict {
			if len(word) > i + 1 {
				continue
			}

			if word == s[i-len(word)+1:i+1] {
				dp[i + 1] = dp[i - len(word) + 1]
				if dp[i + 1] == true {break}
			}
		}
	}
	return dp[len(s)]
}
