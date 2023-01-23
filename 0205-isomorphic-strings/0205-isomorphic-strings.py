class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        st, ts = {}, {}
        for i in range(len(s)):
            st[s[i]] = t[i]
            ts[t[i]] = s[i]
        
        new_s, new_t = "", ""
        for i in range(len(s)):
            new_s += ts[t[i]]
            new_t += st[s[i]]
        
        return s == new_s and t == new_t 
        