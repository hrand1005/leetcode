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
                
            
        