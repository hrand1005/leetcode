/*
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    
    negExp := false
    if n < 0 {
        n = -n
        negExp = true
    }
    
    result := float64(1)
    for n > 0 {
        scale := 1
        multiplier := x
        for scale * 2 < n {
            scale *= 2
            multiplier *= multiplier
        }
        
        result *= multiplier
        n -= scale
    }
    
    if negExp {
        return 1 / result
    }
    return result
}
*/

func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    
    if n < 0 {
        return 1 / myPow(x, -n)
    }
    
    if n % 2 != 0 {
        return x * myPow(x, n-1)
    }
    
    return myPow(x*x, n/2)
}