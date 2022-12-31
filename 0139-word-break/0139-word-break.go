func wordBreak(s string, wordDict []string) bool {
    wordSet := toSet(wordDict)
    seen := make(map[string]bool)
    queue := []string{s}
    
    for len(queue) != 0 {
        word := queue[0]
        queue = queue[1:]
        
        if wordSet[word] || len(word) == 0 {
            return true
        }
        
        for i := 1; i < len(word)+1; i++ {
            if wordSet[word[:i]] && !seen[word[i:]] {
                queue = append(queue, word[i:])
                seen[word[i:]] = true
            }
        }
    }
    
    return false
}

func toSet(s []string) map[string]bool {
    set := make(map[string]bool, len(s))
    for _, v := range s {
        set[v] = true
    }
    return set
}