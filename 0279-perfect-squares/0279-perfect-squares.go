func numSquares(n int) int {
    table := make([]int, n+1)
    table[1] = 1
    for i := 1; i < n+1; i++ {
        minSquares := i
        for j := 1; j < int(math.Pow(float64(i), 0.5)) + 1; j++ {
            minSquares = min(minSquares, 1 + table[i-int(math.Pow(float64(j), 2.0))])
        }
        table[i] = minSquares
    }
    return table[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}