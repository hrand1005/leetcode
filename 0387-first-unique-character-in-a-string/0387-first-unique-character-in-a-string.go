func firstUniqChar(s string) int {
    count := map[rune]int{}

    for _, v := range s {
        count[v]++
    }

    for i, v := range s {
        if count[v] == 1 {
            return i
        }
    }

    return -1
}