func fractionToDecimal(numerator int, denominator int) string {
    result := quotientSign(numerator, denominator)
    numerator, denominator = abs(numerator), abs(denominator)
    quotient, remainder := divmod(numerator, denominator)
    result += strconv.Itoa(quotient)
    
    if remainder == 0 {
        return result
    }
    
    decimal := ""
    seen := map[int]int{ remainder: 0 }
    for remainder != 0 {
        numerator = 10 * remainder
        quotient, remainder = divmod(numerator, denominator)
        decimal += strconv.Itoa(quotient)
        
        if _, ok := seen[remainder]; ok {
            break
        }
        seen[remainder] = len(decimal)
    }
    
    if remainder != 0 {
        repStart := seen[remainder]
        decimal = fmt.Sprintf("%s(%s)", decimal[:repStart], decimal[repStart:])
    }
    
    return fmt.Sprintf("%s.%s", result, decimal) 
}

func quotientSign(dividend, divisor int) string {
    if dividend < 0 && 0 < divisor ||
    divisor < 0 && 0 < dividend {
        return "-"
    } 
    return ""
}

func divmod(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}