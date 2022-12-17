func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    
    if len(matrix) == 1 {
        return matrix[0]
    }
    
    border, matrix := popBorder(matrix)
    return append(border, spiralOrder(matrix)...)
}

func popBorder(matrix [][]int) ([]int, [][]int) {
    // pop the top row
    result := matrix[0]
    matrix = matrix[1:]
    if len(matrix[0]) == 0 {
        return result, matrix
    }
    
    // pop the right column top-down
    rightIdx := len(matrix[0]) - 1
    for i := 0; i < len(matrix) - 1; i++ {
        elem := matrix[i][rightIdx]
        matrix[i] = matrix[i][:rightIdx]
        result = append(result, elem)
    }
    
    // pop the bottom row and reverse
    bottom := matrix[len(matrix)-1]
    matrix = matrix[:len(matrix)-1]
    for i := len(bottom)-1; i >= 0; i-- {
        result = append(result, bottom[i])
    }
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return result, matrix
    }
    
    // pop the left column bottom-up
    for i := len(matrix) - 1; i >= 0; i-- {
        elem := matrix[i][0]
        matrix[i] = matrix[i][1:]
        result = append(result, elem)
    }
    
    return result, matrix
}