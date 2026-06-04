func canJump(nums []int) bool {
    // [1,2,0,1,0] 
	// nums[i] = you can jump till index -> i + nums[i]
	// we have to tell if we can reach last index starting from 0
	// we can iterate l->r and keep record maxidx where i can reach
	// maxidx = max(maxidx, i + nums[i])
	// if ever i > maxidx return false


	maxIdx := 0
	for i, val := range nums {
		if i > maxIdx {
			return false
		}

		maxIdx = max(maxIdx, i + val)
	}
	return true
}
