/*
func isPowerOfThree(n int) bool {
    if n == 1 {
        return true
    }

    if n < 1 || n % 3 != 0 {
        return false
    } 

    return isPowerOfThree(n/3)
}

func isPowerOfThree(n int) bool {
    if n < 1 {
        return false
    }

    for n != 1 {
        if n % 3 != 0 {
            return false
        }
        n /= 3
    }

    return true
}
*/

import (
    "math"
)

func isPowerOfThree(n int) bool {
    if n < 1 {
        return false
    }

    // dividing log(n) by log(3) gives us log_3(n)
    x := math.Round(math.Log(float64(n)) / math.Log(3))

    return n == int(math.Pow(3, x))
}