class Solution {
    /**
     * @param {string} s1
     * @param {string} s2
     * @return {boolean}
     */


    checkInclusion(s1, s2) {
        if (s1.length > s2.length) return false;

        let need = Array(26).fill(0);
        let window = Array(26).fill(0);

        for (let ch of s1) {
            need[ch.charCodeAt(0) - 97]++;
        }

        let k = s1.length;

        for (let i = 0; i < s2.length; i++) {
            window[s2.charCodeAt(i) - 97]++;

            if (i >= k) {
                window[s2.charCodeAt(i - k) - 97]--;
            }

            if (arraysEqual(need, window)) {
                return true;
            }
        }

        return false;
    }
}

function arraysEqual(a, b) {
    for (let i = 0; i < 26; i++) {
        if (a[i] !== b[i]) return false;
    }
    return true;
}