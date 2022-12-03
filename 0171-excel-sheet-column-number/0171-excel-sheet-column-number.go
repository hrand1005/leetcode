/*
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

func recursiveSolution(charMap map[rune]int, multiplier int, columnTitle string) int {
    if len(columnTitle) == 0 {
        return 0
    }:
    lastChar := rune(columnTitle[len(columnTitle)-1])
    thisVal := int(math.Pow(float64(26), float64(multiplier))) * charMap[lastChar]
    return thisVal + recursiveSolution(charMap, multiplier+1, columnTitle[:len(columnTitle)-1])
}

func titleToNumber(columnTitle string) int {
    charMap := initCharMap()
    return recursiveSolution(charMap, 0, columnTitle)
}
*/

func titleToNumber(columnTitle string) int {
    columnNum := 0
    for _, v := range columnTitle {
        columnNum *= 26
        columnNum += int(v - 'A') + 1
    }
    
    return columnNum
}