class Solution {
public:
    int maxArea(vector<int>& heights) {
        int l = 0, r = heights.size() - 1;
        int ans = 0;
        while(l < r) {
            int water = min(heights[l], heights[r]) * (r - l);
            ans = max(ans, water);
            if (heights[r] < heights[l]) r--;
            else l++;
        }
        return ans;
    }
};
