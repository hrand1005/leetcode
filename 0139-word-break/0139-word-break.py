class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        queue = [s]
        word_set = set(wordDict)
        seen = set()
        
        while queue:
            word = queue.pop(0)
            if word in word_set or len(word) == 0:
                return True
            
            for i in range(1, len(word)+1):
                if word[:i] in word_set and word[i:] not in seen:
                    queue.append(word[i:])
                    seen.add(word[i:])
                    
        return False        