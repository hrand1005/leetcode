func largestAltitude(gain []int) int {
    height, highest := 0, 0
    for _, g := range gain {
        height += g
        highest = max(highest, height)
    }
    return highest
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}