import (
    "strings"
)

/*
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

func countAndSay(n int) string {
    say := map[int]string{
        1: "1",
    }
    
    for i := 2; i < n+1; i++ {
        sayThis := say[i-1]
        newSay := ""
        count := 1
        current := string(sayThis[0])
        
        for j := 1; j < len(sayThis); j++ {
            if sayThis[j] == sayThis[j-1] {
                count++
            } else {
                newSay += fmt.Sprintf("%v%v", count, current)
                current = string(sayThis[j])
                count = 1
            }
        }
        newSay += fmt.Sprintf("%v%v", count, current)
        say[i] = newSay
    }
    
    return say[n]
}
*/

func countAndSay(n int) string {
    return countAndSayRecursive(n-1, "1")
}

func countAndSayRecursive(n int, say string) string {
    if n == 0 {
        return say
    }
    
    count := 1
    current := string(say[0])
    newSay := &strings.Builder{}
    
    for i := 1; i < len(say); i++ {
        if say[i] == say[i-1] {
            count++
        } else {
            newSay.WriteString(fmt.Sprintf("%v%v", count, current))
            current = string(say[i])
            count = 1
        }
    }
    newSay.WriteString(fmt.Sprintf("%v%v", count, current))
    
    return countAndSayRecursive(n-1, newSay.String())
}