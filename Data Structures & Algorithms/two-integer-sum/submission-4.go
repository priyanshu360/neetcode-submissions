func twoSum(nums []int, target int) []int {
    i, j := 0, len(nums) - 1
    oldnums := make([]int, len(nums))
    copy(oldnums, nums)

    sort.Ints(nums)

    for ; target != nums[i] + nums[j]; {
        if nums[i] + nums[j] > target {
            j = j - 1
        }else{
            i = i + 1
        }
    }

    val1 := nums[i]
    val2 := nums[j]

    i = -1
    j = -1
    for idx, val := range oldnums {
        if val == val1 && i == -1 {
            i = idx
        }else if val == val2 && j == -1 {
            j = idx
        }
    }

    if i > j {
        t := i
        i = j
        j = t
    }

    return []int{i, j}
}
