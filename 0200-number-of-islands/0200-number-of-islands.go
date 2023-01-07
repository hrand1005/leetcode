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
        neighbors := getNeighbors(cur[0], cur[1], grid)
        for _, n := range neighbors {
            if string(grid[n[0]][n[1]]) == "1" {
                queue = append(queue, n)
            }
        }
    }
    return
}

func getNeighbors(i, j int, grid[][]byte) [][]int {
    neighbors := make([][]int, 0, 4)
    
    if 0 < i {
        neighbors = append(neighbors, []int{i-1, j})
    }
    if i < len(grid)-1 {
        neighbors = append(neighbors, []int{i+1, j})
    }
    if 0 < j {
        neighbors = append(neighbors, []int{i, j-1})
    }
    if j < len(grid[i])-1 {
        neighbors = append(neighbors, []int{i, j+1})
    }
    
    return neighbors
}
