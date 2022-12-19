/*
func setZeroes(matrix [][]int)  {
    zeroRows := map[int]bool{}
    zeroCols := map[int]bool{}
    
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                zeroRows[i] = true
                zeroCols[j] = true
            }
        }
    }
    
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if zeroRows[i] || zeroCols[j] {
                matrix[i][j] = 0
            }
        }
    }
    
    return
}
*/

func setZeroes(matrix [][]int)  {
    firstRowZero := false
    for _, v := range matrix[0] {
        if v == 0 {
            firstRowZero = true
            break
        }
    }
    
    for i := 1; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }
    
    for i := 1; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            for j := 0; j < len(matrix[i]); j++ {
                matrix[i][j] = 0
            }
        }
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    
    if firstRowZero {
        for i := 0; i < len(matrix[0]); i++ {
            matrix[0][i] = 0
        }
    }
    
    return
}