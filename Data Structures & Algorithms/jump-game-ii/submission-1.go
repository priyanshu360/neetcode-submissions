func jump(nums []int) int {
    // similar to jump game 1
	// but here we have to return min number of steps 
	// and last index is always reachable

	// let's say from i index we can reach till j so we will always jump to
	// index between i..j from where we can reach till max index 
	// max(nums[i + 1] + i + 1, nums[i + 2] + i + 2, ... nums[i + j] + i + j) 
	

	ans := 0
	canReach := 0
	preReach := 0
	for i := range nums {
		if i == len(nums) - 1 {
			return ans
		}
		canReach = max(canReach, i + nums[i])	
		if i == preReach {
			preReach = canReach
			ans++
		}
	}
	return 0
}
