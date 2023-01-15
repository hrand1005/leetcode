/*
import (
    "math"
)

func numSquares(n int) int {
    table := make([]int, n+1)
    table[1] = 1
    
    for i := 1; i < n+1; i++ {
        minSquares := i
        for j := 1; j < int(math.Pow(float64(i), 0.5)) + 1; j++ {
            minSquares = min(minSquares, 1 + table[i-j*j])
        }
        table[i] = minSquares
    }
    
    return table[n]
}

*/
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func numSquares(n int) int {
    cache := make(map[int]int, n)
    squares := make([]int, 0, n)
    
    for i := 1; i < n; i++ {
        cache[i*i] = 1
        squares = append(squares, i*i)
    }
    
    var numSquaresRecursive func(int) int
    numSquaresRecursive = func(num int) int {
        if count, ok := cache[num]; ok {
            return count
        }
        
        minSquares := num
        for _, s := range squares {
            if num < s {
                break
            }
            minSquares = min(minSquares, 1+numSquaresRecursive(num-s))
        }
        
        cache[num] = minSquares
        return minSquares
    }
    
    return numSquaresRecursive(n)
}
