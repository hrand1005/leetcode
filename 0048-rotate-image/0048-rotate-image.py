"""                        
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        rotated = set()
        
        for i in range(len(matrix)):
            for j in range(i, len(matrix) - i):
                last = matrix[i][j]
                new_i = j
                new_j = len(matrix) - 1 - i
                cur = matrix[new_i][new_j]
                
                
                for k in range(4):
                    if f'{new_i},{new_j}' not in rotated:
                        matrix[new_i][new_j] = last
                        
                        rotated.add(f'{new_i},{new_j}')
                        
                        last = cur

                        temp = new_i
                        new_i = new_j 
                        new_j = len(matrix) - 1 - temp

                        cur = matrix[new_i][new_j]

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        quadrant_height = len(matrix) // 2
        quadrant_width = len(matrix) - (len(matrix) // 2)
        n = len(matrix) - 1
        
        for i in range(quadrant_height):
            for j in range(quadrant_width):
                matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] = \
                matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j] 
        
                
"""

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        matrix.reverse()
        
        for i in range(len(matrix)):
            for j in range(i, len(matrix)):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
                