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
"""

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        length = len(matrix)
        n = length - 1
        
        for i in range(length // 2):
            for j in range(length - length // 2):
                matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] = \
                matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j] 
        
                