func isIsomorphic(s string, t string) bool {
    smap := make(map[string]string)
    tmap := make(map[string]string)
    for i := 0; i < len(s); i++ {
        smap[string(s[i])] = string(t[i])
        tmap[string(t[i])] = string(s[i])
    }
    
    newS := ""
    newT := ""
    for i := 0; i < len(s); i++ {
        newS += tmap[string(t[i])]
        newT += smap[string(s[i])]
    }
    
    return s == newS && t == newT
}