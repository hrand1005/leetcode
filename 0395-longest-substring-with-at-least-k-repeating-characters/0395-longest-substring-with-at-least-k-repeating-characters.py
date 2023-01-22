class Solution:
    def longestSubstring(self, s: str, k: int) -> int:
        self.cache = {}
        return self.longest_sub(s, k)
        
    def longest_sub(self, s: str, k: int) -> int:
        if len(s) < k:
            return 0
        if self.cache.get(s) != None:
            return self.cache.get(s)
        
        candidates = self.get_candidates(s, k)
        if len(candidates) == 0:
            return 0
        
        valid = []
        for c in candidates:
            if self.valid_sub(c, k):
                valid.append(len(c))
                self.cache[c] = len(c)
            else:
                valid.append(self.longestSubstring(c, k))
                
        self.cache[s] = max(valid)
        return max(valid)        
        
        
    def valid_sub(self, s: str, k: int) -> bool:
        occurrences = {}
        for i in range(len(s)):
            occurrences[s[i]] = occurrences.get(s[i], 0) + 1
        for v in occurrences.values():    
            if v < k:
                return False
        return True    
        
    def get_candidates(self, s :str, k: int) -> List[str]:
        occurrences = {}
        for i in range(len(s)):
            occurrences[s[i]] = occurrences.get(s[i], 0) + 1
        
        candidates = []
        candidate = ""
        for i in range(len(s)):
            if occurrences[s[i]] < k:
                if k <= len(candidate):
                    candidates.append(candidate)
                candidate = ""
            else:    
                candidate += s[i]    
            
        if candidate != "":
            candidates.append(candidate)    
        
        return candidates
        
        
        
        