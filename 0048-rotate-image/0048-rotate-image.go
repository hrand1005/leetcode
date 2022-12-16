/*
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
*/

func rotate(matrix [][]int)  {
    quartileHeight := len(matrix) / 2
    quartileWidth := len(matrix) - (len(matrix) / 2)
    end := len(matrix) - 1
    
    for i := 0; i < quartileHeight; i++ {
        for j := 0; j < quartileWidth; j++ {
            matrix[i][j], matrix[j][end-i], matrix[end-i][end-j], matrix[end-j][i] = 
            matrix[end-j][i], matrix[i][j], matrix[j][end-i], matrix[end-i][end-j]
        }
    }
}