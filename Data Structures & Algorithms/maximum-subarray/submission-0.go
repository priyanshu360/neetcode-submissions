func maxSubArray(nums []int) int {
	//  [2,-3, | 4,-2,2,1,-1,4| ] => 8
	// l->r if our sum < 0 then there is no use of continue the same subarray i.e
	// we can break and check for new 
	

	ans := nums[0] 
	temp := 0
	for _, i := range nums {
		ans = max(ans, i) // if all are negative
		if temp + i  < 0 {
			temp = 0
			continue
		}
		temp += i
		ans = max(ans, temp)
	}

	return ans
}
