class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        islands = 0
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == "1":
                    islands += 1
                    self.traverse_island(i, j, grid)
        return islands            
        
    def traverse_island(self, i: int, j: int, grid: List[List[str]]) -> None:
        queue = [(i, j)]
        
        while queue:
            pos_i, pos_j = queue.pop()
            grid[pos_i][pos_j] = "2"
            neighbors = [
                (pos_i+1, pos_j),
                (pos_i-1, pos_j),
                (pos_i, pos_j+1),
                (pos_i, pos_j-1),
            ]
            
            for n in neighbors:
                if 0 <= n[0] < len(grid) and 0 <= n[1] < len(grid[0]):
                    if grid[n[0]][n[1]] == "1":
                        queue.append(n)
        return                