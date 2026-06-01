func longestConsecutive(nums []int) int {
    res := 0
    store := make(map[int]struct{})
    for _, num := range nums {
        store[num] = struct{}{}
    }

    for _, num := range nums {
        streak, curr := 0, num
        for _, ok := store[curr]; ok; _, ok = store[curr] {
            streak++
            curr++
        }
        if streak > res {
            res = streak
        }
    }
    return res
}