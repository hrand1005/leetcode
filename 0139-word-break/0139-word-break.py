class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        queue = [("", s)]
        word_set = set(wordDict)
        seen = set()
        
        while queue:
            cur, st = queue.pop(0)
            if st in word_set or len(st) == 0:
                return True
            
            seen.add(st)
            for i in range(1, len(st)+1):
                if st[:i] in word_set and st[i:] not in seen:
                    word_set.add(cur+st[:i])
                    queue.append((cur+st[:i], st[i:]))
                    
        return False        