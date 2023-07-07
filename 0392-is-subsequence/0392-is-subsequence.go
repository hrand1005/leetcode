func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    
    stack := []rune(s)
    for _, n := range t {
        if n == stack[0] {
            stack = stack[1:]
        }
        if len(stack) == 0 {
            return true
        }
    }
    return false
}