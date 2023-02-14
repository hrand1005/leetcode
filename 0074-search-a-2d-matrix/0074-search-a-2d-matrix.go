func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    low, high := 0, (m * n) - 1
    for low <= high {
        mid := (low + high) / 2
        row := mid / n
        col := mid % n
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return false
}