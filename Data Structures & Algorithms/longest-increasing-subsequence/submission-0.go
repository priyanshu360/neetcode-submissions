func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums) + 1) 

	const INF = 1_000_000_009
	for i := 0; i <= len(nums); i++ {
		dp[i] = INF
	}

	dp[0] = -INF

	bs := func (num int) int {
		l, r := 0, len(dp) + 1

		for r - l > 1 {
			mid := (l + r) /  2 

			if dp[mid] < num {
				l = mid
			}else {
				r = mid
			}
		}
		return l
	}

	ans := 0
	for _, num := range nums {
		idx := bs(num)
		dp[idx + 1] = min(dp[idx + 1], num)
		ans = max(ans, idx + 1)
	}
	return ans
}
