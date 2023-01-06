import (
    "sort"
)

func largestNumber(nums []int) string {
    strs := make([]string, len(nums))
    for i := 0; i < len(strs); i++ {
        strs[i] = strconv.Itoa(nums[i])
    }
    
    sort.Slice(strs, func(i, j int) bool {
        return strs[j] + strs[i] < strs[i] + strs[j]
    })
    
    if strs[0] == "0" {
        return "0"
    }
    
    return strings.Join(strs, "")
}