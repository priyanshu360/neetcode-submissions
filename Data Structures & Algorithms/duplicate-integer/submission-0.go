func hasDuplicate(nums []int) bool {
    have := make(map[int]bool)
    for _, val := range nums {
        if _, ok := have[val]; ok {
            return true
        }
        have[val] = true;
    }
    return false
}
