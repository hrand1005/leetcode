class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        result = []
        anagram = {}
        
        for i in range(len(strs)):
            this = strs[i]
            sortedStr = ''.join(sorted(this))
            
            group = anagram.get(sortedStr)
            if group != None:
                result[group].append(this)
            else:
                result.append([this])
                anagram[sortedStr] = len(result) - 1
            
        return result    
