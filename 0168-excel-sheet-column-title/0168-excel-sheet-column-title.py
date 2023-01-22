class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        title = ""
        while 0 < columnNumber:
            rem = columnNumber % 26 or 26
            title = str(self.to_letter(rem)) + title
            columnNumber = (columnNumber - rem) // 26
        return title    

    def to_letter(self, x: int) -> str:
        return chr(x+64)
