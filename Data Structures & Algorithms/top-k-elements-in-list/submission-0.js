class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
    const freq = {};

    nums.forEach(n => freq[n] = (freq[n] || 0) + 1)

    return Object.keys(freq)
        .sort((a, b) => freq[b] - freq[a])
        .slice(0, k)
        .map(Number);
    }
}
