func recursiveSolution(n, oneBefore, twoBefore int) int {
    if n == 0 {
        return twoBefore + oneBefore
    }
    
    return recursiveSolution(n - 1, twoBefore + oneBefore, oneBefore)
}


func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    
    return recursiveSolution(n-3, 2, 1)
}

/*
// NOTE: Extremely Slow!
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    
    return climbStairs(n-1) + climbStairs(n-2)
}
*/

/*
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    
    cur := 3
    oneBefore := 2
    twoBefore := 1
    
    for i := 3; i <= n; i++ {
        cur = oneBefore + twoBefore
        twoBefore = oneBefore
        oneBefore = cur
    }
    
    return cur
}
*/

/*
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    
    stepPerms := make([]int, n+1)
    stepPerms[1] = 1
    stepPerms[2] = 2
    
    for i := 3; i <= n; i++ {
        stepPerms[i] = stepPerms[i-1] + stepPerms[i-2]
    }
    
    return stepPerms[n]
}
*/
