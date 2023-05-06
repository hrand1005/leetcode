class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(9):
            if not self.valid_row(board, i) or \
                not self.valid_col(board, i) or \
                not self.valid_square(board, i):
                return False
        return True    
    
    def valid_row(self, board: List[List[str]], r: int) -> bool:
        seen = set()
        for v in board[r]:
            if v in seen:
                return False
            if v != ".":
                seen.add(v)
        return True
    
        
    def valid_col(self, board: List[List[str]], c: int) -> bool:
        seen = set()
        for row in board:
            v = row[c]
            if v in seen:
                return False
            if v != ".":
                seen.add(v)
        return True
    
    def valid_square(self, board: List[List[str]], s: int) -> bool:
        r = (s // 3) * 3
        c = (s % 3) * 3
        seen = set()
        for i in range(r, r+3):
            for j in range(c, c+3):
                v = board[i][j]
                if v in seen:
                    return False
                if v != ".":
                    seen.add(v)
        return True