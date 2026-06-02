func canPartition(nums []int) bool {
	dp := make([]bool, 1_000_01)
	dp[0] = true
	total := 0
	for _, num := range nums {
		total += num
		for sum := 1_000_00; sum >= 0; sum-- {
			if num > sum { continue }

			dp[sum] = dp[sum] || dp[sum - num]
		}
	}

	if total % 2 != 0 {
		return false
	}

	return dp[total/2]
}
