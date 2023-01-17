const (
    willDie = -1
    willLive = 2
)

func gameOfLife(board [][]int)  {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            n := neighbors(i, j, board)
            if board[i][j] == 0 && n == 3 {
                board[i][j] = willLive
            }
            if board[i][j] == 1 {
                if n < 2 || 3 < n {
                    board[i][j] = willDie
                }
            }
        }
    }
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == willDie {
                board[i][j] = 0
            }
            if board[i][j] == willLive {
                board[i][j] = 1
            }
        }
    }
}

func neighbors(i, j int, board [][]int) int {
    neighborPositions := [][]int{
        {i-1, j-1}, {i-1, j}, {i-1, j+1},
        {i+1, j-1}, {i+1, j}, {i+1, j+1},
        {i, j-1}, {i, j+1},
    }
    
    neighbors := 0
    for _, n := range neighborPositions {
        posI, posJ := n[0], n[1]
        if 0 <= posI && posI < len(board) && 0 <= posJ && posJ < len(board[posI]) {
            if board[posI][posJ] == 1 || board[posI][posJ] == willDie {
                neighbors++
            }
        }
    }
    return neighbors
}