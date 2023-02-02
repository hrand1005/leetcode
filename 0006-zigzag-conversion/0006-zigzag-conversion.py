class Solution:
    def convert(self, s: str, numRows: int) -> str:
        rows = {}
        chars = list(s)
        
        while chars:
            for i in range(numRows):
                if chars:
                    rows[i] = rows.get(i, "") + chars.pop(0)
            for i in range(numRows-2, 0, -1):
                if chars:
                    rows[i] = rows.get(i, "") + chars.pop(0)
            
        zigzag = ""
        for line in rows.values():
            zigzag += line
            
        return zigzag    