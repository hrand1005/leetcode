func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    
    table := make([][]int, m)
    for i := 0; i < m; i++ {
        table[i] = make([]int, n)
    }
    
    for i := 1; i < m; i++ {
        table[i][0] = table[i-1][0] + grid[i-1][0]
    }
    for j := 1; j < n; j++ {
        table[0][j] = table[0][j-1] + grid[0][j-1]
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            top := table[i-1][j] + grid[i-1][j]
            left := table[i][j-1] + grid[i][j-1]
            table[i][j] = min(top, left)
        }
    }
    
    return table[m-1][n-1] + grid[m-1][n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}