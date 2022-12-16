func rotate(matrix [][]int)  {
    end := len(matrix) - 1
    
    for i := 0; i < len(matrix) / 2; i++ {
        matrix[i], matrix[end-i] = matrix[end-i], matrix[i]
    }
    
    for i := 0; i < len(matrix); i++ {
        for j := i; j < len(matrix); j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
}