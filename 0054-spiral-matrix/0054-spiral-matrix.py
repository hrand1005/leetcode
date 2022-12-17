class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        if len(matrix) == 0:
            return []
        if len(matrix) == 1:
            return matrix[0]
        
        result, inner_matrix = self.pop_border(matrix) 
        return result + self.spiralOrder(inner_matrix)
        
    def pop_border(self, matrix: List[List[int]]) -> (List, List[List[int]]):
        border = matrix.pop(0)
        
        if len(matrix[0]) == 0:
            return border, [[]]
        
        for i in range(0, len(matrix)-1):
            right_elem = matrix[i].pop(len(matrix[i])-1)
            border.append(right_elem)
        
        reverse_bottom = matrix.pop(len(matrix)-1)[::-1]
        border.extend(reverse_bottom)
        
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return border, matrix
        
        for j in range(len(matrix)-1, 0, -1):
            left_elem = matrix[j].pop(0)
            border.append(left_elem)
        
        return border, matrix  