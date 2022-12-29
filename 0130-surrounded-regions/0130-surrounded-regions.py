"""
class Solution:
    def solve(self, board: List[List[str]]) -> None:
        path = set()
        for i in range(len(board)):
            for j in range(len(board[i])):
                if board[i][j] == 'O' and board[i][j] not in path:
                    board, new_path = self.update_region(board, (i, j))
                    path = path.union(new_path)
        
        return board
    
    def update_region(self, board: List[List[str]], start: tuple[int, int]) -> (List[List[str]], Set[tuple[int, int]]):
        path = [start]
        queue = [start]
        visited = set()
        flip = True
        
        while queue:
            pos_i, pos_j = queue.pop(0)
            visited.add((pos_i, pos_j))
            
            if len(board) - 1 < pos_i or pos_i < 0:
                flip = False
                break
            if len(board[0]) - 1 < pos_j or pos_j < 0:
                flip = False
                break
            
            if board[pos_i][pos_j] == 'X':
                continue
            else:
                path.append((pos_i, pos_j))
            
            neighbors = [
                (pos_i+1, pos_j),
                (pos_i-1, pos_j),
                (pos_i, pos_j+1),
                (pos_i, pos_j-1),
            ]
            
            for n in neighbors:
                if n not in visited:
                    queue.append(n)
            
        if flip:
            for pos in path:
                board[pos[0]][pos[1]] = 'X'

        return board, path        
"""
    
class Solution:
    def solve(self, board: List[List[str]]) -> None:
        # scan for starting elements on the outside ring of the matrix
        queue = self.get_border_entries(board)
        visited = set()
        
        while queue:
            i, j = queue.pop(0)
            visited.add((i, j))
            
            neighbors = [
                (i+1, j),
                (i-1, j),
                (i, j+1),
                (i, j-1),
            ]
            
            for n in neighbors:
                if 0 <= n[0] <= len(board)-1 and 0 <= n[1] <= len(board[0])-1:
                    if board[n[0]][n[1]] == 'O' and n not in visited:
                        queue.append(n)
        
        for i in range(len(board)):
            for j in range(len(board[i])):
                if board[i][j] == 'O' and (i, j) not in visited:
                    board[i][j] = 'X'
        
    def get_border_entries(self, board: List[List[str]]) -> List[tuple[int, int]]:
        starts = []
        # scan top
        for i in range(len(board[0])):
            if board[0][i] == 'O':
                starts.append((0, i))
                
        # scan left        
        for i in range(len(board)):
            if board[i][0] == 'O':
                starts.append((i, 0))
        
        # scan right
        right = len(board[0]) - 1
        for i in range(len(board)):
            if board[i][right] == 'O':
                starts.append((i, right))
                
        # scan bottom
        bottom = len(board) - 1
        for i in range(len(board[-1])):
            if board[bottom][i] == 'O':
                starts.append((bottom, i))
                
        return starts        
