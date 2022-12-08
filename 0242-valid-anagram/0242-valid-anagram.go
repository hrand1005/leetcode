func isAnagram(s string, t string) bool {
    occurrences := map[rune]int{}

    for _, v := range s {
        occurrences[v]++
    }

    for _, v := range t {
        occurrences[v]--
    }

    for _, v := range occurrences {
        if v != 0 {
            return false
        }
    }

    return true
}