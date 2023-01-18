/*
import (
    "math"
)

func coinChange(coins []int, amount int) int {
    table := make([]int, amount+1)
    for i := 1; i < len(table); i++ {
        minCoins := math.MaxInt32
        for _, c := range coins {
            if i-c >= 0 && table[i-c] != math.MaxInt32 {
                minCoins = min(minCoins, 1+table[i-c])
            }
        }
        table[i] = minCoins
    }
    if table[amount] == math.MaxInt32 {
        return -1
    }
    return table[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
*/

func coinChange(coins []int, amount int) int {
    cache := make(map[int]int)
    
    var coinChangeRecursive func([]int, int) int
    coinChangeRecursive = func(coins []int, amount int) int {
        if v, ok := cache[amount]; ok {
            return v
        }
        if amount == 0 {
            return 0
        }
        minCoins := math.MaxInt32
        for _, c := range coins {
            if amount-c >= 0 {
                minCoins = min(minCoins, 1+coinChangeRecursive(coins, amount-c))
            }
        }
        cache[amount] = minCoins
        return minCoins
    }
    
    res := coinChangeRecursive(coins, amount)
    if res == math.MaxInt32 {
        return -1
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}