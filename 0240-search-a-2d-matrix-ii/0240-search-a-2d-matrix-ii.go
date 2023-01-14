/*
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
*/

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    
    topRight := matrix[0][len(matrix[0])-1]
    
    if topRight == target {
        return true
    }
    
    if topRight < target {
        return searchMatrix(matrix[1:], target)
    }
    
    return searchMatrix(removeRightCol(matrix), target)
}

func removeRightCol(matrix [][]int) [][]int {
    newMatrix := make([][]int, len(matrix))
    for i, row := range matrix {
        newMatrix[i] = row[:len(row)-1]
    }
    return newMatrix
}