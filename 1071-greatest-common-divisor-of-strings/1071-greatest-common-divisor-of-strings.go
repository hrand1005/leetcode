func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }
    gcd := calculateGCD(len(str1), len(str2))
    return str1[:gcd]
}

func calculateGCD(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

/*
func gcdOfStrings(str1 string, str2 string) string {
    i := 0
    sub := ""
    maxSub := ""
    minLen := min(len(str1), len(str2))
    for i < minLen && str1[i] == str2[i] {
        sub += string(str1[i])
        if dividesString(str1, sub) && dividesString(str2, sub) {
            maxSub = sub
        }
        i++
    }
    
    return maxSub
}

func dividesString(s, sub string) bool {
    return strings.Repeat(sub, len(s)/len(sub)) == s
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
*/