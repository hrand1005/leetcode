/*
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
*/

type parenCombo struct {
    str string
    left int
    right int
}

func generateParenthesis(n int) []string {
    allCombos := []string{}
    
    stack := []*parenCombo{
        {
            str: "(",
            left: n - 1,
            right: n,
        },
    }
    
    for len(stack) > 0 {
        popped := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if popped.left > popped.right {
            continue
        }
        
        if popped.left == 0 && popped.right == 0 {
            allCombos = append(allCombos, popped.str)
        }
        
        if popped.left > 0 {
            newCombo := &parenCombo{
                str: popped.str + "(",
                left: popped.left-1,
                right: popped.right,
            }
            stack = append(stack, newCombo)
        }
        
        if popped.right > 0 {
            newCombo := &parenCombo{
                str: popped.str + ")",
                left: popped.left,
                right: popped.right-1,
            }
            stack = append(stack, newCombo)
        }
    }
    
    return allCombos
}