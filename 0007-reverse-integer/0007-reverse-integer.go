import (
    "math"
)

func reverse(x int) int {
    result := 0
    for x != 0 {
        digit := x % 10
        if willOverflow(result, digit) {
            return 0
        }
        result = result * 10 + digit
        x /= 10
    }
    
    return result 
}

func willOverflow(a, b int) bool {
    if a < math.MinInt32 / 10 || math.MaxInt32 / 10 < a {
        return true
    } else if a == math.MinInt32 / 10 {
        return b < -8
    } else if a == math.MaxInt32 / 10 {
        return b > 7
    }
    return false
}
