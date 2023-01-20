class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        row = [1]
        for i in range(rowIndex):
            new_row = []
            for i in range(1, len(row)):
                num = row[i-1] + row[i]
                new_row.append(num)
            row = [1] + new_row + [1]
        return row    
                    