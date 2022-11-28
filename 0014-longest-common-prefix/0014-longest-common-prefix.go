import (
    "strings"
)

func longestCommonPrefix(strs []string) string {
    var longestPrefix strings.Builder
    for i := 0; i < len(strs[0]); i++ {
        ch := strs[0][i]
        for _, s := range strs[1:] {
            if i >= len(s) || s[i] != ch {
                return longestPrefix.String()
            }
        }
        longestPrefix.WriteByte(ch)
    }
    
    return longestPrefix.String()
}