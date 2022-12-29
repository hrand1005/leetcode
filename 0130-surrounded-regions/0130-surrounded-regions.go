type point struct {
    i, j int
}

func (p point) String() string {
    return fmt.Sprintf("%v,%v", p.i, p.j)
}

func solve(board [][]byte)  {
    queue := getBorderPoints(board)
    visited := make(map[string]bool)
    
    for len(queue) > 0 {
        pos := queue[0]
        queue = queue[1:]
        visited[pos.String()] = true
        
        neighbors := []point{
            {pos.i-1, pos.j},
            {pos.i+1, pos.j},
            {pos.i, pos.j-1},
            {pos.i, pos.j+1},
        }
        
        for _, n := range neighbors {
            if 0 <= n.i && n.i < len(board)-1 && 0 <= n.j && n.j <= len(board[0])-1 {
                if board[n.i][n.j] == 'O' && !visited[n.String()] {
                    queue = append(queue, n)
                }
            }
        }
    }
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            pos := point{i, j}
            if board[i][j] == 'O' && !visited[pos.String()] {
                board[i][j] = 'X'
            }
        }
    }
}

func getBorderPoints(board [][]byte) []point {
    entryPoints := make([]point, 0, 2*len(board) + 2*len(board[0]))
    // scan top
    for i, v := range board[0] {
        if v == 'O' {
            entryPoints = append(entryPoints, point{ 0, i })
        }
    }
    
    // scan left
    for i, v := range board {
        if v[0] == 'O' {
            entryPoints = append(entryPoints, point{ i, 0 })
        }
    }
    
    // scan right
    for i, v := range board {
        if v[len(v)-1] == 'O' {
            entryPoints = append(entryPoints, point{ i, len(v)-1 })
        }
    }
    
    // scan bottom
    for i, v := range board[len(board)-1] {
        if v == 'O' {
            entryPoints = append(entryPoints, point{ len(board)-1, i })
        }
    }
    
    return entryPoints
}