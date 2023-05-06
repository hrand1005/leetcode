func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        if !validRow(board, i) || !validCol(board, i) || !validSquare(board, i) {
            return false
        }
    }
    return true
}

func validRow(board [][]byte, r int) bool {
    seen := make(map[byte]bool)
    for _, v := range board[r] {
        if v != '.' && seen[v] {
            return false
        }
        seen[v] = true
    }
    return true
}

func validCol(board [][]byte, c int) bool {
    seen := make(map[byte]bool)
    for _, row := range board {
        v := row[c]
        if v != '.' && seen[v] {
            return false
        }
        seen[v] = true
    }
    return true
}

func validSquare(board [][]byte, s int) bool {
    row := (s / 3) * 3
    col := (s % 3) * 3
    seen := make(map[byte]bool)
    for i := row; i < row+3; i++ {
        for j := col; j < col+3; j++ {
            v := board[i][j]
            if v != '.' && seen[v] {
                return false
            }
            seen[v] = true
        }
    }
    return true
}