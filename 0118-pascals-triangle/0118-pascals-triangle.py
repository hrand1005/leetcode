"""
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        above = [1]
        result = [above]
        for i in range(1, numRows):
            new_row = []
            for j in range(i+1):
                if j == 0 or j == i:
                    new_row.append(1)
                else:
                    new_row.append(above[j]+above[j-1])
                    
            result.append(new_row)
            above = new_row
            
        return result
"""
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        result = [[1]*n for n in range(1, numRows+1)]
        for i in range(1, numRows):
            for j in range(1, i):
                result[i][j] = result[i-1][j-1] + result[i-1][j]
                
        return result