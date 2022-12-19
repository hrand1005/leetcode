"""
# O(m + n)
class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        zero_rows = set()
        zero_cols = set()
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[i][j] == 0:
                    zero_rows.add(i)
                    zero_cols.add(j)
                    
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if i in zero_rows or j in zero_cols:
                    matrix[i][j] = 0
        
        return
"""

class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        first_row_zero = 0 in matrix[0]
        
        for i in range(1, len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[i][j] == 0:
                    matrix[0][j] = 0
                    matrix[i][0] = 0
        
        for i in range(1, len(matrix)):
            if 0 in matrix[i]:
                matrix[i] = [0] * len(matrix[i])
            else:    
                for j in range(0, len(matrix[i])):
                    print(f'(0, j): (0, {j})')
                    if matrix[0][j] == 0:
                        print(f'setting ({i}, {j}) to zero')
                        matrix[i][j] = 0
                        
        if first_row_zero:
            matrix[0] = [0] * len(matrix[0])
        
        return