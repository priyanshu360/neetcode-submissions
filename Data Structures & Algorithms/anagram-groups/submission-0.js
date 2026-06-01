class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
    const map = {};
    for (let s of strs) {
        const key = [...s].sort().join("");
        (map[key] ??= []).push(s);
    }
    return Object.values(map);
    }

}
