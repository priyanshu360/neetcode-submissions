class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        map <string, vector<string>> mp;

        for(auto str : strs) {
            auto ns = str;
            sort(ns.begin(), ns.end());

            if (mp.find(ns) == mp.end()) {
                mp[ns] = vector<string> {};
            }
            mp[ns].push_back(str);
        }

        vector <vector<string>> result;
        for(auto it = mp.begin(); it != mp.end(); it++) {
            vector <string> value = (*it).second;

            result.push_back(value);
        }

        return result;
    }
};
