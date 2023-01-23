import (
    "strings"
)

func isIsomorphic(s string, t string) bool {
    smap := make(map[byte]byte)
    tmap := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        smap[s[i]] = t[i]
        tmap[t[i]] = s[i]
    }
    
    var newS strings.Builder
    var newT strings.Builder
    for i := 0; i < len(s); i++ {
        newS.WriteByte(tmap[t[i]])
        newT.WriteByte(smap[s[i]])
    }
    
    return s == newS.String() && t == newT.String()
}