func generateParenthesis(n int) []string {
    return generate(n, n, "", []string{})
}

func generate(left, right int, current string, all []string) []string {
    if left > right {
        return nil
    }
    
    if left == 0 && right == 0 {
        all = append(all, current)
        return all
    }
    
    lResult, rResult := []string{}, []string{}
    if left > 0 {
        lResult = generate(left - 1, right, current + "(", all)
    }
    
    if right > 0 {
        rResult = generate(left, right - 1, current + ")", all)
    }
    
    return append(lResult, rResult...)
}