func strStr(haystack string, needle string) int {
    for i, v := range haystack {
        if v == rune(needle[0]) {
            if i + len(needle) > len(haystack) {
                return -1
            }
            
            candidate := haystack[i:i+len(needle)]
            if candidate == needle {
                return i
            }
        }
    }
    
    return -1
}