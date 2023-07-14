func removeStars(s string) string {
    stack := []rune{}
    for _, l := range s {
        if l == '*' {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, l)
        }
    }
    return string(stack)
}