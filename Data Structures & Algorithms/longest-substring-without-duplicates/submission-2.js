class Solution {
    /**
     * @param {string} s
     * @return {number}
     */
    lengthOfLongestSubstring(s) {
        let fzArray = Array(1000).fill(0);
        let k = 1;

        let l = 0, r = 0, ans = 0;

        while (r < s.length) {
            fzArray[s.charCodeAt(r)]++;
            let ok = Math.max(...fzArray) <= k;

            while (!ok) {
                fzArray[s.charCodeAt(l)]--;
                l++;

                ok = Math.max(...fzArray) <= k;
            }

            ans = Math.max(ans, r - l + 1);
            r++;
        }

        return ans;
    }
}
