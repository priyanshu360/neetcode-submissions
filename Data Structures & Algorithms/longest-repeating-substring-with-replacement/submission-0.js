class Solution {
    /**
     * @param {string} s
     * @param {number} k
     * @return {number}
     */
    characterReplacement(s, k) {
    let fzArray = Array(26).fill(0);

    let l = 0, r = 0, ans = 0;

    while (r < s.length) {
        fzArray[s.charCodeAt(r) - 65]++;

        let ok =
            (fzArray.reduce((sum, v) => sum + v, 0) - Math.max(...fzArray)) <= k;

        while (!ok) {
            fzArray[s.charCodeAt(l) - 65]--;
            l++;

            ok =
                (fzArray.reduce((sum, v) => sum + v, 0) - Math.max(...fzArray)) <= k;
        }

        ans = Math.max(ans, r - l + 1);
        r++;
    }

    return ans;
}



    // L = -1; R = N

    // R - L > 1 {
    //     if ond(M) : L = M
    //     else R = M 
    // }

    // [F .... F F] last true L or first false R

    // L = -1;  R = N
    // R - L > 1 {
    //     if cond(M) : R = M
    //     else L = M 
    // }
    // [F F T .... T T T] first T -> R

}
