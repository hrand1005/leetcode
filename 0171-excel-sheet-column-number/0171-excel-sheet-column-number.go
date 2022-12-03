import (
    "math"
)

func initCharMap() map[rune]int {
    const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    charMap := map[rune]int{}
    for i := 0; i < 26; i++ {
        thisChar := rune(letters[i])
        charMap[thisChar] = i + 1
    }
    
    return charMap
}

func titleToNumber(columnTitle string) int {
    charMap := initCharMap()
    colNum := 0
    multiplier := 0
    
    for i := len(columnTitle)-1; i >= 0; i-- {
        thisChar := rune(columnTitle[i])
        thisVal := charMap[thisChar]
        colNum += int(math.Pow(float64(26), float64(multiplier))) * thisVal
        multiplier++
    }
    
    return colNum
}