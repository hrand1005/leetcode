func reverseWords(s string) string {
    raw := strings.Split(s, " ")
    
    words := make([]string, 0, len(raw))
    for _, w := range raw {
        if w != "" {
            words = append(words, w)
        }
    }
    
    for i := 0; i < len(words)/2; i++ {
        words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
    }
    return strings.Join(words, " ")
}