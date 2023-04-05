/*
func fib(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}

func fib(n int) int {
    memo := map[int]int{
        0: 0,
        1: 1,
    }
    var fibRecurse func(int) int
    fibRecurse = func(n int) int {
        if v, ok := memo[n]; ok {
            return v
        }
        memo[n] = fib(n-1) + fib(n-2)
        return memo[n]
    }
    return fibRecurse(n)
}
*/

func fib(n int) int {
    if n <= 1 {
        return n
    }
    twoBefore, oneBefore := 0, 1
    for i := 2; i < n; i++ {
        twoBefore, oneBefore = oneBefore, oneBefore + twoBefore
    }
    return twoBefore + oneBefore
}