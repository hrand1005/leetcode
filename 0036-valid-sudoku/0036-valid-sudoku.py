class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(9):
            if not self.valid_row(board, i):
                return False
            if not self.valid_column(board, i):
                return False
            if not self.valid_square(board, i):
                return False
            
        return True     
            
    def valid_row(self, board: List[List[str]], index: int) -> bool:
        row = board[index]
        occurred = {}
        for n in row:
            if n != '.' and occurred.get(n) != None:
                print(f'Encountered {n} again in row')
                return False
            occurred[n] = 1
        
        return True
        
    def valid_column(self, board: List[List[str]], index: int) -> bool:
        occurred = {}
        for i in range(9):
            n = board[i][index]
            if n != '.' and occurred.get(n) != None:
                print(f'Encountered {n} again in column')
                return False
            occurred[n] = 1
        
        return True
        
    def valid_square(self, board: List[List[str]], index: int) -> bool:
        i_offset = (index % 3) * 3
        j_offset = (index // 3) * 3
        
        occurred = {}
        for i in range(3):
            for j in range(3):
                n = board[i_offset+i][j_offset+j]
                if n != '.' and occurred.get(n) != None:
                    print(f'Encountered {n} again in square')
                    return False
                occurred[n] = 1
                
        return True       
        