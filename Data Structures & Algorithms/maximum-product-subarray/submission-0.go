func maxProduct(nums []int) int {
	maxProd := nums[0]
	minProd := nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if num < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(num, maxProd*num)
		minProd = min(num, minProd*num)

		ans = max(ans, maxProd)
	}

	return ans
}
