class Solution {
    /**
     * @param {number[]} temperatures
     * @return {number[]}
     */
    dailyTemperatures(temperatures) {
        let stack = []
        let ans = new Array(temperatures.length).fill(0)

        for (let i = 0; i < temperatures.length; i++) {
            
            while(stack.length && 
            temperatures[stack[stack.length - 1]] < temperatures[i]) {
                let idx = stack.pop()
                ans[idx] = i - idx
            }
            
            stack.push(i)
        }

        return ans;

    }
}
