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
"""

class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        maxHeap = []
        for row in matrix:
            for elem in row:
                heappush(maxHeap, -elem)
                if k < len(maxHeap):
                    heappop(maxHeap)
        return -heappop(maxHeap)            

"""
class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        minHeap = []
        for row in matrix:
            for elem in row:
                
"""