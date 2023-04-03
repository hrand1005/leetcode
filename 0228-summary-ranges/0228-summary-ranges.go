func summaryRanges(nums []int) []string {
    if len(nums) < 1 {
        return nil
    }
    ranges := make([][]int, 0, len(nums))
    start := nums[0]
    current := []int{ start, start }
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] + 1 {
            current[1] = nums[i]
        } else {
            ranges = append(ranges, current)
            current = []int{nums[i], nums[i]}
        }
    }
    ranges = append(ranges, current)
    return formatRanges(ranges)
}

func formatRanges(ranges [][]int) []string {
    rangeStrings := make([]string, len(ranges))
    for i := 0; i < len(rangeStrings); i++ {
        if ranges[i][0] == ranges[i][1] {
            rangeStrings[i] = fmt.Sprintf("%v", ranges[i][0])
        } else {
            rangeStrings[i] = fmt.Sprintf("%v->%v", ranges[i][0], ranges[i][1])
        }
    }
    return rangeStrings
}