class Solution:
    def partition(self, s: str) -> List[List[str]]:
        if len(s) == 0:
            return [[]]
        
        if len(s) == 1:
            return [[s]]
        
        all_pals = []
        for i in range(len(s), 0, -1):
            candidate = s[:i]
            if self.is_palindrome(candidate):
                remainders = self.partition(s[i:])
                for r in remainders:
                    all_pals.append([candidate]+r)
        
        return all_pals
                
    def is_palindrome(self, s: str) -> bool:
        if len(s) <= 1:
            return True
        if s[0] == s[-1]:
            return self.is_palindrome(s[1:-1])
        return False