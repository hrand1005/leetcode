"""
class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        stack = [ (0, 0) ]
        finish = (m-1, n-1)
        paths = 0
        while stack:
            last = stack.pop()
            if last == finish:
                paths += 1
                continue
            
            neighbors = self.get_neighbors(last, m, n)
            stack.extend(neighbors)
            
        return paths    
    
    def get_neighbors(self, node: tuple[int, int], m: int, n: int) -> List[tuple[int, int]]:
        neighbors = []
        if node[0] < m - 1:
            neighbors.append((node[0]+1, node[1]))
            
        if node[1] < n - 1:
            neighbors.append((node[0], node[1]+1))
        
        return neighbors

class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        table = [ [1] * n for i in range(m) ]
        for i in range(1, m):
            for j in range(1, n):
                table[i][j] = table[i-1][j] + table[i][j-1]
        
        return table[-1][-1]
"""

class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        current_line = [1] * n
        
        for i in range(1, m):
            for j in range(1, n):
                current_line[j] += current_line[j-1]
        
        return current_line[-1]