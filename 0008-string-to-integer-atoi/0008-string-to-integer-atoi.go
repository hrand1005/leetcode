import (
    "fmt"
    "math"
    "strings"
)

var runeToInt = map[rune]int{
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
}

func myAtoi(s string) int {
    s = strings.TrimLeft(s, " ")
    
    if len(s) == 0 {
        return 0
    }
    
    multiplier := 1
    if s[0] == '-' {
        multiplier = -1
        s = s[1:]
    } else if s[0] == '+' {
        s = s[1:]
    }
    
    result := 0
    for len(s) > 0 {
        v, ok := runeToInt[rune(s[0])]
        if !ok {
            break
        }
        
        limit, ok := willOverflow(result * multiplier, v)
        if ok {
            return limit
        }
        
        result = result * 10 + v
        s = s[1:]
    }
    
    return result * multiplier
}

// willOverflow returns whether the int v will 
// overflow, and which limit it will exceed
func willOverflow(res, add int) (int, bool) {
    if res > math.MaxInt32 / 10 + 7 - add {
        return math.MaxInt32, true 
    } 
    if res < math.MinInt32 / 10 - 8 + add{
        return math.MinInt32, true 
    }
    
    return 0, false
}