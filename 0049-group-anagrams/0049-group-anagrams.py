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