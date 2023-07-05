func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := 0
    for _, v := range candies {
        if v > maxCandies {
            maxCandies = v
        }
    }
    
    result := make([]bool, len(candies))
    for i, v := range candies {
        result[i] = v + extraCandies >= maxCandies
    }
    return result
}