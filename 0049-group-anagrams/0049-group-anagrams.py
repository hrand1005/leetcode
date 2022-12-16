class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        result = []
        anagram = {}
        
        for i in range(len(strs)):
            this = strs[i]
            listStr = list(this)
            listStr.sort()
            sortedStr = ''.join(listStr)
            
            if anagram.get(sortedStr) != None:
                result[anagram.get(sortedStr)].append(this)
            else:
                result.append([this])
                anagram[sortedStr] = len(result) - 1
            
        return result    
        