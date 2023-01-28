/*
func isMatch(s string, p string) bool {
    table := make([][]bool, len(s)+1)
    for i := 0; i < len(table); i++ {
        table[i] = make([]bool, len(p)+1)
    }
    table[0][0] = true
    
    for j := 2; j < len(p)+1; j++ {
        if p[j-1] == '*' {
            table[0][j] = table[0][j-2]
        }
    }
    
    for i := 1; i < len(s)+1; i++ {
        for j := 1; j < len(p)+1; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '.' {
                table[i][j] = table[i-1][j-1]
            } else if p[j-1] == '*' {
                empty := table[i][j-2]
                nonempty := (s[i-1] == p[j-2] || p[j-2] == '.') && table[i-1][j]
                table[i][j] = empty || nonempty
            }
        }
    }
    
    return table[len(s)][len(p)]
}
*/

type pair struct {
    s, p string
}

func isMatch(s string, p string) bool {
    cache := map[pair]bool{
        pair{"", ""}: true,
    }
    
    var matches func(s, p string) bool
    matches = func(s, p string) bool {
        if cache[pair{s, p}] {
            return true
        }
        if p == "" {
            return false
        }
        
        match := false
        if p[len(p)-1] == '*' {
            match = match || matches(s, p[:len(p)-2])
            idx := len(s)-1
            for 0 <= idx && (s[idx] == p[len(p)-2] || p[len(p)-2] == '.') {
                match = match || matches(s[:idx], p[:len(p)-2])
                idx--
            }
        } else if s != "" && (s[len(s)-1] == p[len(p)-1] || p[len(p)-1] == '.') {
            match = match || matches(s[:len(s)-1], p[:len(p)-1])
        }
        
        cache[pair{s, p}] = match
        return match
    }
    
    return matches(s, p)
}