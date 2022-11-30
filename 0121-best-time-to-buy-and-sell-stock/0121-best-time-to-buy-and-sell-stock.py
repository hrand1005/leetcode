class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        min_buy = prices[0]
        for i in range(len(prices)):
            min_buy = min(min_buy, prices[i])
            max_profit = max(max_profit, prices[i] - min_buy)
                    
        return max_profit
