"""
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        self.cache = {}
        return self.coin_change_recursive(coins, amount)

    def coin_change_recursive(self, coins: List[int], amount: int) -> int:
        if self.cache.get(amount) != None:
            return self.cache[amount]
        if amount in coins:
            return 1
        if amount == 0:
            return 0
        
        global_min = -1
        for i in range(len(coins)-1, -1, -1):
            if 0 < amount - coins[i]:
                min_coins = 1 + self.coinChange(coins[:i+1], amount - coins[i])
                if min_coins != 0:
                    if global_min == -1:
                        global_min = min_coins
                    else:    
                        global_min = min(global_min, min_coins)
                    
        self.cache[amount] = global_min
        return global_min            
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