import (
    "strings"
)

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    
    numToLetters := map[string]string{
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
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