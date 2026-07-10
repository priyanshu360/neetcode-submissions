class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map <int, int> valToIdx;

        for(int i = 0; i < nums.size(); i++){
            if(valToIdx.find(target - nums[i]) != valToIdx.end()) return vector<int> {valToIdx[target - nums[i]], i};
            if(valToIdx.find(nums[i]) != valToIdx.end()) continue;
            valToIdx[nums[i]] = i;
        }
        return vector<int>{};

    }
};
