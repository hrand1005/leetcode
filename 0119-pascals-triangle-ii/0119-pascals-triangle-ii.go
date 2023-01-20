func getRow(rowIndex int) []int {
    row := []int{1}
    for i := 0; i < rowIndex; i++ {
        next := make([]int, len(row)+1)
        for j := 1; j < len(row); j++ {
            next[j] = row[j-1] + row[j]
        }
        next[0], next[len(next)-1] = 1, 1
        row = next
    }
    return row
}