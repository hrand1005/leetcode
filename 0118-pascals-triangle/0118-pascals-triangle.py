class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        if numRows == 0:
            return
        
        row1 = [1]
        result = [row1]
        above = row1
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
        