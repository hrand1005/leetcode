func longestSubstring(s string, k int) int {
    charCount := make(map[string]int, len(s))
    for _, c := range s {
        charCount[string(c)]++
    }
    
    splitChar := ""
    for c, v := range charCount {
        if v < k {
            splitChar = c
            break
        }
    }
    
    longest := 0
    if splitChar != "" {
        subs := strings.Split(s, splitChar)
        for _, sub := range subs {
            longest = max(longest, longestSubstring(sub, k))
        }
    } else {
        longest = len(s)
    }
    
    return longest
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}