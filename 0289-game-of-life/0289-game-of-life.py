class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        for i in range(len(board)):
            for j in range(len(board[i])):
                n = self.live_neighbors(i, j, board)
                if board[i][j] == 0 and n == 3:
                    board[i][j] = 2
                elif board[i][j] == 1:
                    if n < 2 or 3 < n:
                        board[i][j] = -1
                        
        for i in range(len(board)):
            for j in range(len(board[i])):
                if board[i][j] == 2:
                    board[i][j] = 1
                elif board[i][j] == -1:
                    board[i][j] = 0
                    
    def live_neighbors(self, i: int, j: int, board: List[List[int]]) -> int:
        positions = [
            (i-1, j-1),
            (i-1, j),
            (i-1, j+1),
            (i, j-1),
            (i, j+1),
            (i+1, j-1),
            (i+1, j),
            (i+1, j+1),
        ]
        neighbors = 0
        for p in positions:
            pos_i, pos_j = p[0], p[1]
            if 0 <= pos_i < len(board) and 0 <= pos_j < len(board[0]):
                if board[pos_i][pos_j] == 1 or board[pos_i][pos_j] == -1:
                    neighbors += 1
        return neighbors        
            
        