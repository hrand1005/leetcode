/*
func uniquePaths(m int, n int) int {
    table := createTable(m, n)
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            table[i][j] = table[i-1][j] + table[i][j-1]
        }
    }
    
    return table[m-1][n-1]
}

func createTable(m, n int) [][]int {
    row := make([]int, n)
    for i := 0; i < n; i++ {
        row[i] = 1
    }
    
    table := make([][]int, m)
    for i := 0; i < m; i++ {
        table[i] = row
    }
    
    return table
}
*/

func uniquePaths(m int, n int) int {
    row := make([]int, n)
    for i := 0; i < n; i++ {
        row[i] = 1
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            row[j] += row[j-1]
        }
    }
    
    return row[n-1]
}