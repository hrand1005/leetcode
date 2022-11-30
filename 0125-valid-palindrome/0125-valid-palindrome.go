import (
    "unicode"
)

func isPalindrome(s string) bool {
    i := 0
    j := len(s) - 1
    
    for i < j {
        left := rune(s[i])
        right := rune(s[j])
        
        if !unicode.IsLetter(left) && !unicode.IsNumber(left) {
            i++
            continue
        }
        if !unicode.IsLetter(right) && !unicode.IsNumber(right) {
            j--
            continue
        }
        if unicode.ToLower(left) != unicode.ToLower(right) {
            return false
        }
        
        i++; j--
    }
    
    return true
}