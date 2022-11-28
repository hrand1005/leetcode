func isValid(s string) bool {
    openToClose := map[rune]rune{
        '{': '}',
        '(': ')',
        '[': ']',
    }
    
    stack := make([]rune, 0, len(s))
    for _, c := range s {
        if _, ok := openToClose[c]; ok {
            stack = append(stack, c)
        } else if len(stack) == 0 || openToClose[stack[len(stack)-1]] != c  {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}