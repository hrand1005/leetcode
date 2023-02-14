import (
    "strings"
)

var numToLetters = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    
    if len(digits) == 1 {
        return strings.Split(numToLetters[digits], "")
    }
    
    prefixes := strings.Split(numToLetters[string(digits[0])], "")
    suffixes := letterCombinations(digits[1:])
    combos := make([]string, 0)
    for _, p := range prefixes {
        for _, s := range suffixes {
            combos = append(combos, p+s)
        }
    }
    return combos
}

/*
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    
    combos := map[int][]string{}
    firstKey := string(digits[0])
    combos[0] = strings.Split(numToLetters[firstKey], "")
    
    for i := 1; i < len(digits); i++ {
        key := string(digits[i])
        for _, prefix := range combos[i-1] {
            newLetters := strings.Split(numToLetters[key], "")
            for _, letter := range newLetters {
                combos[i] = append(combos[i], prefix + string(letter))
            }
        }
    }
    
    return combos[len(digits)-1]
}
*/

/*
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    
    if len(digits) == 1 {
        key := string(digits[0])
        return strings.Split(numToLetters[key], "")
    }
    
    prev := letterCombinations(digits[:len(digits)-1])
    key := string(digits[len(digits)-1])
    newLetters := strings.Split(numToLetters[key], "")
    combos := make([]string, 0, len(prev) * len(newLetters))
    
    for _, prefix := range prev {
        for _, letter := range newLetters {
            combos = append(combos, string(prefix) + string(letter))
        }
    }
    
    return combos
}
*/