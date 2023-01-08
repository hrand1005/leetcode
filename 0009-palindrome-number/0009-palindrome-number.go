func isPalindrome(x int) bool {
    xStr := strconv.Itoa(x)
    return palindrome(xStr)
}

func palindrome(s string) bool {
    if len(s) <= 1 {
        return true
    }
    if s[0] != s[len(s)-1] {
        return false
    }
    return palindrome(s[1:len(s)-1])
}