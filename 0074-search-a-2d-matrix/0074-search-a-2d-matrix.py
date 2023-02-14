NOT_FOUND = 2 ** 32

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        row = self.find_row(matrix, target)
        if row == NOT_FOUND:
            return False
        
        low, high = 0, len(matrix[row])-1
        while low <= high:
            mid = (low + high) // 2
            if matrix[row][mid] == target:
                return True
            elif matrix[row][mid] < target:
                low = mid + 1
            else:
                high = mid - 1
        return False        
                
        
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
        