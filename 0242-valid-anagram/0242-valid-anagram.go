func isAnagram(s string, t string) bool {
    count := make(map[rune]int, len(s))
    for _, l := range s {
        count[l]++
    }
    for _, l := range t {
        count[l]--
    }
    for _, v := range count {
        if v != 0 {
            return false
        }
    }
    return true
}