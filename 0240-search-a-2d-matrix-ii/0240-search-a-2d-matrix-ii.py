"""
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for row in range(len(matrix)):
            if self.found(matrix[row], target, 0, len(matrix[row])-1):
                return True
        return False    
    
    def found(self, row: int, target: int, low: int, high: int) -> bool:
        if high < low:
            return False
        midpoint = (low + high) // 2
        if row[midpoint] == target:
            return True
        if row[midpoint] < target:
            return self.found(row, target, midpoint+1, high)
        if target < row[midpoint]:
            return self.found(row, target, low, midpoint-1)
"""

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for row in matrix:
            if target in row:
                return True
        return False    
