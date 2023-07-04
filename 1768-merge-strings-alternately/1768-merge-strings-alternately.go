func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func mergeAlternately(word1 string, word2 string) string {
    result := ""
    for i := 0; i < max(len(word1), len(word2)); i++ {
        if i < len(word1) {
            result += string(word1[i])
        }
        if i < len(word2) {
            result += string(word2[i])
        }
    }
    return result
}