class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        min_buy = prices[0]
        for i in range(len(prices)):
            if prices[i] < min_buy:
                min_buy = prices[i]
            if prices[i] - min_buy > max_profit:
                max_profit = prices[i] - min_buy
                    
        return max_profit