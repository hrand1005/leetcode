import (
    "sort"
    "strings"
)

func groupAnagrams(strs []string) [][]string {
    result := [][]string{}
    groupToIndex := map[string]int{}
    
    for _, s := range strs {
        slice := strings.Split(s, "")
        sort.Strings(slice)
        anagram := strings.Join(slice, "")
        
        if idx, ok := groupToIndex[anagram]; ok {
            result[idx] = append(result[idx], s)
        } else {
            result = append(result, []string{s})
            groupToIndex[anagram] = len(result) - 1
        }
    }
    
    return result
}