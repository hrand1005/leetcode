import (
    "math"
)

func checkOverflow(sign, quotient, toAdd int) (bool, int) {
    if sign > 0 && quotient > math.MaxInt32 - toAdd  {
        return true, math.MaxInt32
    }
    if sign < 0 && 0 - quotient < math.MinInt32 + toAdd {
        return true, math.MinInt32
    }
    
    return false, 0
}

/*
func divide(dividend int, divisor int) int {
    sign := 1
    if divisor < 0 {
        divisor = -divisor
        sign = -sign
    }
    
    if dividend < 0 {
        dividend = -dividend
        sign = -sign
    }
    
    quotient := 0
    counter := divisor
    
    for counter <= dividend {
        multiplied := 1
        for (counter << 1) <= dividend {
            multiplied <<= 1
            counter <<= 1
        }
        
        overflowed, result := checkOverflow(sign, quotient, multiplied)
        if overflowed {
            return result
        }
        
        dividend -= counter
        quotient += multiplied
        counter = divisor
    }
    
    if sign < 0 {
        return 0 - quotient
    }
    
    return quotient
}
*/

func divide(dividend int, divisor int) int {
    sign := 1
    if divisor < 0 {
        divisor = -divisor
        sign = -sign
    }
    
    if dividend < 0 {
        dividend = -dividend
        sign = -sign
    }
    
    quotient := 0
    counter := divisor
    
    for counter <= dividend {
        multiplied := 1
        for (counter + counter) <= dividend {
            multiplied += multiplied
            counter += counter
        }
        
        overflowed, result := checkOverflow(sign, quotient, multiplied)
        if overflowed {
            return result
        }
        
        dividend -= counter
        quotient += multiplied
        counter = divisor
    }
    
    if sign < 0 {
        return 0 - quotient
    }
    
    return quotient
}