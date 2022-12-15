import (
    "strings"
)

func countAndSay(n int) string {
    say := map[int]string{
        1: "1",
    }
    
    for i := 2; i < n+1; i++ {
        sayThis := say[i-1]
        newSay := &strings.Builder{}
        count := 1
        current := string(sayThis[0])
        
        for j := 1; j < len(sayThis); j++ {
            if sayThis[j] == sayThis[j-1] {
                count++
            } else {
                newSay.WriteString(fmt.Sprintf("%v%v", count, current))
                current = string(sayThis[j])
                count = 1
            }
        }
        newSay.WriteString(fmt.Sprintf("%v%v", count, current))
        say[i] = newSay.String()
    }
    
    return say[n]
}