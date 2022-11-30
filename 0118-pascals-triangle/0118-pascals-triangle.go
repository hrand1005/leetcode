/*
func generate(numRows int) [][]int {
    result := initializeSlice(numRows)
    for i := 1; i < numRows; i++ {
        for j := 1; j < i; j++ {
            result[i][j] = result[i-1][j-1] + result[i-1][j]
        }
    }
    
    return result
}

func initializeSlice(rows int) [][]int {
    s := make([][]int, rows)
    for i := 0; i < rows; i++ {
        for j := 0; j < i+1; j++ {
            s[i] = append(s[i], 1)
        }
    }
    return s
}
*/

func generate(numRows int) [][]int {
    result := [][]int{}
    above := []int{1}
    for i := 0; i < numRows; i++ {
        newRow := []int{}
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                newRow = append(newRow, 1)
            } else {
                newRow = append(newRow, above[j] + above[j-1])
            }
        }
        result = append(result, newRow)
        above = newRow
    }
    return result
}