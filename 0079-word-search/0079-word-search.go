func exist(board [][]byte, word string) bool {
    if !possiblePath(board, word) {
        return false
    }
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if dfs(board, word, i, j) {
                return true
            }
        }
    }
    
    return false
}

func possiblePath(board [][]byte, word string) bool {
    letterCount := map[byte]int{}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            letterCount[board[i][j]]++
        }
    }
    
    for i := range word {
        if letterCount[word[i]] == 0 {
            return false
        }
        letterCount[word[i]]--
    }
    
    return true
}

func dfs(board [][]byte, word string, i, j int) bool {
    if len(word) == 0 {
        return true
    }
    if i < 0 || len(board) - 1 < i {
        return false
    }
    if j < 0 || len(board[i]) - 1 < j {
        return false
    }
    if board[i][j] != word[0] {
        return false
    }
    
    board[i][j] = 0
    
    if dfs(board, word[1:], i-1, j) || dfs(board, word[1:], i+1, j) ||
    dfs(board, word[1:], i, j-1) || dfs(board, word[1:], i, j+1) {
        return true
    }
    
    board[i][j] = word[0]
    
    return false
}