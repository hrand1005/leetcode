func searchMatrix(matrix [][]int, target int) bool {
    row, col := 0, len(matrix[0])-1
    for row <= len(matrix)-1 && 0 <= col {
        if matrix[row][col] == target {
            return true
        } else if target < matrix[row][col] {
            col--
        } else {
            row++
        }
    }
    return false
}