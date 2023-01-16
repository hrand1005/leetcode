func addBinary(a string, b string) string {
    binsum := ""
    carry := 0
    for a != "" || b != "" {
        digitA := 0
        if a != "" {
            digitA, _ = strconv.Atoi(string(a[len(a)-1]))
            a = a[:len(a)-1]
        }
        
        digitB := 0
        if b != "" {
            digitB, _ = strconv.Atoi(string(b[len(b)-1]))
            b = b[:len(b)-1]
        }
        
        digit := digitA + digitB + carry
        carry = 0
        if 1 < digit {
            digit -= 2
            carry = 1
        }
        
        binsum = strconv.Itoa(digit) + binsum
    }
    
    if carry == 1 {
        return "1" + binsum
    }
    return binsum
}