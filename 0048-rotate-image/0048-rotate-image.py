class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        rotated = set()
        
        for i in range(len(matrix)):
            for j in range(i, len(matrix)-i):
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
                        
                        
        
        