class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {string}
     */
    minWindow(s, t) {
        if (t.length > s.length) return "";

        let need = new Map();
        for (let ch of t) {
            need.set(ch, (need.get(ch) || 0) + 1);
        }

        let have = 0;
        let required = need.size;

        let window = new Map();
        let res = [-1, -1];
        let resLen = Infinity;

        let l = 0;

        for (let r = 0; r < s.length; r++) {
            let c = s[r];
            window.set(c, (window.get(c) || 0) + 1);

            if (need.has(c) && window.get(c) === need.get(c)) {
                have++;
            }

            while (have === required) {
                if ((r - l + 1) < resLen) {
                    res = [l, r];
                    resLen = r - l + 1;
                }

                window.set(s[l], window.get(s[l]) - 1);

                if (need.has(s[l]) && window.get(s[l]) < need.get(s[l])) {
                    have--;
                }

                l++;
            }
        }

        return resLen === Infinity
            ? ""
            : s.slice(res[0], res[1] + 1);
    }
}
