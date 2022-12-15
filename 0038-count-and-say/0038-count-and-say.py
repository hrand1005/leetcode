"""
class Solution:
    def countAndSay(self, n: int) -> str:
        return self.countAndSayRecursive(n, 1, "1")
        
    def countAndSayRecursive(self, n: int, i: int, say: str) -> str:
        if i == n:
            return say
        
        new_say = ""
        cur = say[0]
        count = 1
        for j in range(1, len(say)):
            if say[j] == say[j-1]:
                count += 1
            else:
                new_say += str(count) + cur
                count = 1
                cur = say[j]
        
        return self.countAndSayRecursive(n, i+1, new_say + str(count) + cur)
                
"""
class Solution:
    def countAndSay(self, n: int) -> str:
        say = {
            1: "1",
        }
        for i in range(2, n+1):
            this = say[i-1]
            cur = this[0]
            count = 1
            new = ""
            for j in range(1, len(this)):
                if this[j] == this[j-1]:
                    count += 1
                else:
                    new += f'{count}{cur}'
                    count = 1
                    cur = this[j]
            say[i] = new + f'{count}{cur}'
        
        return say[n]
            