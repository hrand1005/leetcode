func isValid(s string) bool {
    match := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := make([]rune, 0, len(s))
    for _, c := range s {
        if _, ok := match[c]; !ok {
            stack = append(stack, c)
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != match[c] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}