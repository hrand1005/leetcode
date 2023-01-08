/*
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
*/

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    before := x
    after := 0
    for 0 < x {
        after *= 10
        after += x % 10
        x /= 10
    }
    
    return before == after
}