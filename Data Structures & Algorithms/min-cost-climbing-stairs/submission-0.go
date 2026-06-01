const INF = 1_000_000_000
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp := make([]int, n + 1)

    for i := 1; i <= n; i++ {
        dp[i] = INF 
    }

    for i := 1; i <= n; i++ {
        dp[i] = min(dp[i], cost[i - 1] + dp[i - 1])

        if i > 1 {
            dp[i] = min(dp[i], cost[i - 1] + dp[i - 2])
        }
    }

    return min(dp[n], dp[n - 1])
}
