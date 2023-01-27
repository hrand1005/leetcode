class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        table = [[False] * (len(p)+1) for _ in range(len(s)+1)]
        table[0][0] = True
        
        for j in range(2, len(p)+1):
            if p[j-1] == "*":
                table[0][j] = table[0][j-2]
        
        for i in range(1, len(s)+1):
            for j in range(1, len(p)+1):
                if s[i-1] == p[j-1] or p[j-1] == ".":
                    table[i][j] = table[i-1][j-1]
                elif p[j-1] == "*":
                    empty = table[i][j-2]
                    nonempty = table[i-1][j] and (s[i-1] == p[j-2] or p[j-2] == ".")
                    table[i][j] = empty or nonempty
                    
        return table[-1][-1]