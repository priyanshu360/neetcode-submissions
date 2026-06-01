class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices) {
        let minPrice = Infinity;

        return prices.reduce((maxProfit, price) => {
            minPrice = Math.min(minPrice, price);
            return Math.max(maxProfit, price - minPrice);
        }, 0);
    }
};
