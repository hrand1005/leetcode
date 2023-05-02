class Solution(object):
    def isAnagram(self, s, t):
        occ = {}
        for l in s:
            if occ.get(l) is None:
                occ[l] = 1
            else:
                occ[l] += 1
        
        for l in t:
            if occ.get(l, 0) == 0:
                return False
            occ[l] -= 1
            
        return len(s) == len(t)