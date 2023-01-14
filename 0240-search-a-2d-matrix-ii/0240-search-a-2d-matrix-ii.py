"""

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for row in matrix:
            if target in row:
                return True
        return False    
"""
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        row, col = 0, len(matrix[0])-1
        while row <= len(matrix)-1 and 0 <= col:
            if matrix[row][col] == target:
                return True
            elif target < matrix[row][col]:
                col -= 1
            else:
                row += 1
        return False        