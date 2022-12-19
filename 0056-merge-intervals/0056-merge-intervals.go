import (
    "sort"
)

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    result := [][]int{ intervals[0] }
    
    for _, v := range intervals[1:] {
        if v[0] <= result[len(result)-1][1] {
            result[len(result)-1][1] = max(result[len(result)-1][1], v[1])
        } else {
            result = append(result, v)
        }
    }
    
    return result
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}