func numIslands(grid [][]byte) int {
    islands := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if string(grid[i][j]) == "1" {
                traverseIsland(i, j, grid)
                islands++
            }
        }
    }
    return islands
}

func traverseIsland(i, j int, grid [][]byte) {
    queue := [][]int{
        {i, j},
    }
    for len(queue) != 0 {
        cur := queue[len(queue)-1]
        queue = queue[:len(queue)-1]
        
        grid[cur[0]][cur[1]] = 0
        neighbors := [][]int{
            {cur[0]-1, cur[1]},
            {cur[0]+1, cur[1]},
            {cur[0], cur[1]-1},
            {cur[0], cur[1]+1},
        }
        
        for _, n := range neighbors {
            if 0 <= n[0] && n[0] < len(grid) && 0 <= n[1] && n[1] < len(grid[0]) {
                if string(grid[n[0]][n[1]]) == "1" {
                    queue = append(queue, n)
                }
            }
        }
    }
    return
}