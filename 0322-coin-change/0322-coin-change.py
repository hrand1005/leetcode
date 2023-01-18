MAX_INT = 2**32

class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        self.cache = {0: 0}
        min_coins = self.coin_change_recursive(coins, amount)
        if min_coins == MAX_INT:
            return -1
        return min_coins

    def coin_change_recursive(self, coins: List[int], amount: int) -> int:
        if self.cache.get(amount) != None:
            return self.cache[amount]
        if amount in coins:
            self.cache[amount] = 1
            return 1
        
        min_coins = MAX_INT
        for c in coins:
            if amount-c >= 0:
                min_coins = min(min_coins, 1+self.coin_change_recursive(coins, amount-c))
                
        self.cache[amount] = min_coins
        return min_coins            
"""
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        MAX_INT = float('inf')
        table = [0] + [MAX_INT] * (amount)
        
        for i in range(1, len(table)):
            min_coins = MAX_INT
            for c in coins:
                if i-c == 0:
                    min_coins = min(min_coins, 1)
                elif i-c > 0 and table[i-c] != MAX_INT:
                    min_coins = min(min_coins, 1+table[i-c])
            table[i] = min_coins
        
        if table[amount] == MAX_INT:
            return -1
        return table[amount]
"""