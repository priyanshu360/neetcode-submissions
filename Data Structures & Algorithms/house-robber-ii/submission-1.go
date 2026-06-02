func rob(nums []int) int {
	n := len(nums)
	dp1, dp2 := make([]int, len(nums)), make([]int, len(nums))

	if n == 1 {
		return nums[0]
	}

	for i := 1; i < n; i++ {
		var s1, s2 int
		if i > 1 {
			s1 = dp1[i - 2]
			s2 = dp2[i - 2]
		}
		dp1[i] = max(dp1[i - 1], nums[i - 1] + s1)
		dp2[i] = max(dp2[i - 1], nums[i] + s2)
	}
	return max(dp1[n - 1], dp2[n - 1])
}
