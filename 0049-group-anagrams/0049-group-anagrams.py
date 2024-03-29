"""
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for s in strs:
            key = self.get_hash(s)
            groups[key] = groups.get(key, []) + [s]
        return groups.values()
    
    def get_hash(self, s: str) -> str:
        l = list(s)
        l.sort()
        return "".join(l)
"""

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for s in strs:
            occ = [0] * 26
            for l in s:
                occ[ord(l)-ord("a")] += 1
            key = tuple(occ)    
            groups[key] = groups.get(key, []) + [s]
        return groups.values()    