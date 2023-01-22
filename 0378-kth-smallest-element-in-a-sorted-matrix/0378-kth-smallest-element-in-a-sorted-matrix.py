"""
class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        flat = self.merge_sort(matrix)
        return flat[k-1]
    
    def merge_sort(self, matrix: List[List[int]]) -> List[int]:
        flat_sorted = []
        while matrix:
            row = matrix.pop(0)
            flat_sorted = self.merge(flat_sorted, row)
        return flat_sorted
            
    def merge(self, l1: List[int], l2: List[int]) -> List[int]:
        res = []
        while 0 < len(l1) and 0 < len(l2):
            if l1[0] < l2[0]:
                res.append(l1.pop(0))
            else:
                res.append(l2.pop(0))
        if len(l1) != 0:
            res += l1
        else:
            res += l2
        return res

class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        max_heap = []
        for row in matrix:
            for elem in row:
                heappush(max_heap, -elem)
                if k < len(max_heap):
                    heappop(max_heap)
        return -heappop(max_heap)            

class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        min_heap = []
        for i in range(min(len(matrix), k)):
            heappush(min_heap, (matrix[i][0], i, 0))
        
        kth = 0
        for i in range(k):
            kth, row, col = heappop(min_heap)
            if col + 1 < len(matrix): 
                heappush(min_heap, (matrix[row][col+1], row, col+1))
                
        return kth        
                
"""
class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        low, high = matrix[0][0], matrix[-1][-1]
        kth = 0
        while low <= high:
            mid = (low + high) // 2
            if k <= self.count_leq(mid, matrix):
                kth = mid
                high = mid - 1
            else:
                low = mid + 1
        return kth        
    
    def count_leq(self, mid, matrix):
        count = 0
        col = len(matrix) - 1
        for row in range(len(matrix)):
            while 0 <= col and mid < matrix[row][col]:
                col -= 1
            count += col + 1    
        return count    