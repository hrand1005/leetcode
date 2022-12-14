const blank = byte(46)

func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        if !validRow(board, i) {
            return false
        }
        
        if !validColumn(board, i) {
            return false
        }
        
        if !validSquare(board, i) {
            return false
        }
    }
    
    return true
}

func validRow(board [][]byte, index int) bool {
    seen := map[byte]bool{}
    row := board[index]
    
    for _, v := range row {
        n := v
        if _, ok := seen[n]; ok && n != blank {
            return false
        }
        seen[n] = true
    }
    
    return true
}

func validColumn(board [][]byte, index int) bool {
    seen := map[byte]bool{}
    
    for i := 0; i < 9; i++ {
        n := board[i][index]
        if _, ok := seen[n]; ok && n != blank {
            return false
        }
        seen[n] = true
    }
    
    return true
}

func validSquare(board [][]byte, index int) bool {
    seen := map[byte]bool{}
    
    iOffset := (index % 3) * 3
    jOffset := (index / 3) * 3
    
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            n := board[iOffset + i][jOffset + j]
            if _, ok := seen[n]; ok && n != blank {
                return false
            }
            seen[n] = true
        }
    }
    
    return true
}