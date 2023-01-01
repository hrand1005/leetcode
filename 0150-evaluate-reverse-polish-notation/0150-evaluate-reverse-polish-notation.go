func evalRPN(tokens []string) int {
    stack := []int{}
    for _, t := range tokens {
        if isOperator(t) {
            y, x := stack[len(stack)-1], stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            switch t {
            case "+":
                stack = append(stack, x + y)
            case "-":
                stack = append(stack, x - y)
            case "*":
                stack = append(stack, x * y)
            case "/":
                stack = append(stack, x / y)
            }
        } else {
            i, _ := strconv.Atoi(t)
            stack = append(stack, i)
        }
    }
    return stack[0]
}

func isOperator(s string) bool {
    if s == "+" || s == "-" || s == "*" || s == "/" {
        return true
    }
    return false
}