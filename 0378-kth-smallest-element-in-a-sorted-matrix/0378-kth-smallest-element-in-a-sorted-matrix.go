func kthSmallest(matrix [][]int, k int) int {
    low := matrix[0][0]
    high := matrix[len(matrix)-1][len(matrix)-1]
    
    var kth int
    for low <= high {
        mid := (low + high) / 2
        if k <= countLeq(mid, matrix) {
            kth = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    
    return kth
}

func countLeq(x int, matrix [][]int) int {
    count := 0
    col := len(matrix) - 1
    for _, row := range matrix {
        for 0 <= col && x < row[col] {
            col--
        }
        count += col + 1
    }
    return count
}