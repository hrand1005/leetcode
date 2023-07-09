func equalPairs(grid [][]int) int {
    rows := make(map[string]int)
    for _, row := range grid {
        key := ""
        for _, v := range row {
            key += strconv.Itoa(v) + ","
        }
        rows[key]++
    }
    
    pairs := 0
    for i := 0; i < len(grid); i++ {
        key := ""
        for j := 0; j < len(grid); j++ {
            key += strconv.Itoa(grid[j][i]) + ","
        }
        pairs += rows[key]
    }
    
    return pairs
}