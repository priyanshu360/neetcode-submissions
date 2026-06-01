class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    maxSlidingWindow(nums, k) {
        let deque = []
        let ans = []
        let max = 0
        for (let i = 0; i < nums.length; i++) {
            let val = nums[i]
            while(deque.length > 0  && nums[deque[deque.length - 1]] < val) {
                deque = deque.slice(0, deque.length - 1)
            }
            deque.push(i)
            while(deque.length > 0 && deque[0] <= i - k ){
                deque = deque.slice(1, deque.length)
            }
            max =  nums[deque[0]]
            if (i + 1 < k) {
                continue
            }

            ans.push(max)
        }
        return ans
    }
}
