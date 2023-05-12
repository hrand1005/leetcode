import (
    "strconv"
)

func evalRPN(tokens []string) int {
    stack := make([]int, 0, len(tokens))
    for _, t := range tokens {
        switch t {
        case "+":
            a, b := stack[len(stack)-1], stack[len(stack)-2]
            stack = append(stack[:len(stack)-2], a+b)
        case "-":
            a, b := stack[len(stack)-1], stack[len(stack)-2]
            stack = append(stack[:len(stack)-2], b-a)
        case "/":
            a, b := stack[len(stack)-1], stack[len(stack)-2]
            stack = append(stack[:len(stack)-2], b/a)
        case "*":
            a, b := stack[len(stack)-1], stack[len(stack)-2]
            stack = append(stack[:len(stack)-2], a*b)
        default:
            n, _ := strconv.Atoi(t)
            stack = append(stack, n)
        }
    }
    return stack[len(stack)-1]
}