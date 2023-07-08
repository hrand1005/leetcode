func maxVowels(s string, k int) int {
    vowelCount := 0
    for i := 0; i < k; i++ {
        if isVowel(string(s[i])) {
            vowelCount++
        }
    }
    maxVowels := vowelCount
    
    l, r := 0, k
    for r < len(s) {
        if maxVowels == k {
            return maxVowels
        }
        if isVowel(string(s[l])) {
            vowelCount--
        }
        if isVowel(string(s[r])) {
            vowelCount++
        }
        l++
        r++
        maxVowels = max(maxVowels, vowelCount)
    }
    return maxVowels
}

func isVowel(s string) bool {
    vowels := "aeiou"
    return strings.Contains(vowels, s)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}