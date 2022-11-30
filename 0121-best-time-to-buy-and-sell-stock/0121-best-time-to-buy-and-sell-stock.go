func maxProfit(prices []int) int {
    maxProfit := 0
    minBuy := prices[0]
    for _, v := range prices {
        if v < minBuy {
            minBuy = v
        }
        if v - minBuy > maxProfit {
            maxProfit = v - minBuy
        }
    }
    return maxProfit
}