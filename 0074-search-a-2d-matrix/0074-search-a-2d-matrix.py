"""
NOT_FOUND = 2 ** 32

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        row = self.find_row(matrix, target)
        if row == NOT_FOUND:
            return False
        col = self.find_col(matrix[row], target)
        if col == NOT_FOUND:
            return False
        return True        
                
    def find_row(self, matrix: List[List[int]], target: int) -> int:
        low, high = 0, len(matrix)-1
        while low <= high:
            mid = (low + high) // 2
            if matrix[mid][0] <= target <= matrix[mid][-1]:
                return mid
            elif matrix[mid][-1] < target:
                low = mid + 1
            else:
                high = mid - 1
        return NOT_FOUND            
        
    def find_col(self, row: List[int], target: int) -> int:
        low, high = 0, len(row)-1
        while low <= high:
            mid = (low + high) // 2
            if row[mid] == target:
                return mid
            elif row[mid] < target:
                low = mid + 1
            else:
                high = mid - 1
        return NOT_FOUND        
"""

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        low, high = 0, (m * n) - 1
        while low <= high:
            mid = (low + high) // 2
            row = mid // n
            col = mid % n
            
            if matrix[row][col] == target:
                return True
            elif matrix[row][col] < target:
                low = mid + 1
            else:
                high = mid - 1
        return False