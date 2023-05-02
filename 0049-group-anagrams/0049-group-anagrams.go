func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string, len(strs))
    for _, s := range strs {
        hash := getHash(s)
        groups[hash] = append(groups[hash], s)
    }
    
    result := make([][]string, 0, len(groups))
    for _, g := range groups {
        result = append(result, g)
    }
    return result
}

func getHash(s string) string {
    rs := []rune(s)
    sort.Slice(rs, func(i, j int) bool {
        return rs[i] < rs[j]
    })
    return string(rs)
}