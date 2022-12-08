func lengthOfLongestSubstring(s string) int {
    str := []rune(s)
    maxSub := 0

    for i := 0; i < len(str); i++ {

        seen := map[rune]bool{}
        subStr := []rune{}

        for j := i; j < len(str); j++ {
            if v := seen[str[j]]; v {
                break
            }

            seen[str[j]] = true
            subStr = append(subStr, str[j])
        }

        if len(subStr) > maxSub {
            maxSub = len(subStr)
        }
    }

    return maxSub
}

/*
func lengthOfLongestSubstring(s string) int {
    last := map[rune]int{}

    maxSub := 0
    left := 0

    for i, v := range s {
        if idx, ok := last[v]; !ok || idx < left {
            sub := i - left + 1
            if sub > maxSub {
                maxSub = sub
            }
        } else {
            left = last[v] + 1
        }

        last[v] = i
    }

    return maxSub
}
*/