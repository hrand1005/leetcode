import (
    "strings"
)

func intToRoman(num int) string {
    vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    
    var roman strings.Builder
    for num != 0 {
        for i := 0; i < len(vals); i++ {
            if vals[i] <= num {
                num -= vals[i]
                roman.WriteString(symbols[i])
                break
            }
        }
    }
    
    return roman.String()
}