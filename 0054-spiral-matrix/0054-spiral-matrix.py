class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        if len(matrix) == 0:
            return []
        if len(matrix) == 1:
            return matrix[0]
        
        outer, inner = self.get_outer_inner(matrix) 
        return outer + self.spiralOrder(inner)
        
    def get_outer_inner(self, matrix: List[List[int]]) -> (List, List[List[int]]):
        outer = matrix.pop(0)
        
        if len(matrix[0]) == 0:
            return outer, [[]]
        
        for i in range(0, len(matrix)-1):
            right_elem = matrix[i].pop(len(matrix[i])-1)
            outer.append(right_elem)
        
        reverse_bottom = matrix.pop(len(matrix)-1)[::-1]
        outer.extend(reverse_bottom)
        
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return outer, matrix
        
        for j in range(len(matrix)-1, 0, -1):
            left_elem = matrix[j].pop(0)
            outer.append(left_elem)
        
        return outer, matrix  