func lengthOfLastWord(s string) int {
    last := len(s)-1
    for 0 <= last && s[last] == ' ' {
        last--
    }
    
    count := 0
    for 0 <= last && s[last] != ' ' {
        last--
        count++
    }
    
    return count
}