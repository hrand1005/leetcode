import (
    "math"
)

func reverse(x int) int {
    multiplier := 1
    if x < 0 {
        x *= -1
        multiplier = -1
    }
    
    result := 0
    for x > 0 {
        result = (result * 10) + (x % 10)
        x -= x % 10
        x /= 10
    }
    
    result *= multiplier
    if result < math.MinInt32 || math.MaxInt32 < result {
        return 0
    }
    
    return result 
}