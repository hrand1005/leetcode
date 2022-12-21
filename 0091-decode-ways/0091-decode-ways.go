/*
import (
    "strconv"
)

func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    } 
    if len(s) == 1 {
        return 1
    }
    
    allWays := make([]int, len(s)+1)
    allWays[0] = 1
    
    if s[1] != '0' {
        allWays[1] = 1
    }
    
    for i := 2; i < len(s) + 1; i++ {
        if s[i-1] != '0' {
            allWays[i] += allWays[i-1]
        }
        
        double, _ := strconv.Atoi(string(s[i-2:i]))
        if 10 <= double && double <= 26 {
            allWays[i] += allWays[i-2]
        }
    }
    
    return allWays[len(s)]
}
*/

func numDecodings(s string) int {
	n := len(s)
	m := make(map[string]int, n+1)
    var dfs func(string) int
    dfs = func(cur string) int {
        if len(cur) == 0 {
			return 1
		}
		if cur[0] == '0' {
			return 0
		}
        if len(cur) == 1 {
            return 1
        }
		if count, ok := m[cur]; ok {
			return count
		}
        
        ways := dfs(cur[1:])
        double, _ := strconv.Atoi(string(cur[:2]))
        if 10 <= double && double <= 26 {
            ways += dfs(cur[2:])
        }
        
		m[cur] = ways
		return ways
	}
	return dfs(s)
}