"""
class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        MAX_INT = 2 ** 32
        table = [[MAX_INT] * len(grid[i]) for i in range(len(grid))]
        table[0][0] = 0
        
        for i in range(len(grid)):
            for j in range(0, len(grid[i])):
                if i == 0 and j == 0:
                    table[i][j] = 0
                elif i == 0:    
                    table[i][j] = table[i][j-1] + grid[i][j-1]
                elif j == 0:
                    table[i][j] = table[i-1][j] + grid[i-1][j]
                else:
                    from_top = table[i-1][j] + grid[i-1][j]
                    from_left = table[i][j-1] + grid[i][j-1]
                    table[i][j] = min(from_top, from_left)
                    
        return table[-1][-1] + grid[-1][-1]
"""
class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        MAX_INT = 2 ** 32
        table = [[MAX_INT] * len(grid[i]) for i in range(len(grid))]
        table[0][0] = 0
        
        for i in range(1, len(grid)):
            table[i][0] = table[i-1][0] + grid[i-1][0]
        for j in range(1, len(grid[0])):
            table[0][j] = table[0][j-1] + grid[0][j-1]
        for i in range(1, len(grid)):
            for j in range(1, len(grid[i])):
                top = table[i-1][j] + grid[i-1][j]
                left = table[i][j-1] + grid[i][j-1]
                table[i][j] = min(top, left)
                
        return table[-1][-1] + grid[-1][-1]