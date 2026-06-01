func climbStairs(n int) int {
    dp := make([]int, n + 1)

    dp[0] = 1

    for i := 1; i <= n; i++ {
        if i > 1 {
            dp[i] += dp[i - 2]
        }
        dp[i] += dp[i - 1]
    }

    return dp[n]
}
