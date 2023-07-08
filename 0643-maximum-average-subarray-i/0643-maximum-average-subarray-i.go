func findMaxAverage(nums []int, k int) float64 {
    queue := nums[:k]
    qTotal := sum(queue)
    maxTotal := qTotal
    for _, v := range nums[k:] {
        qTotal = qTotal - queue[0] + v
        queue = append(queue[1:], v)
        maxTotal = max(maxTotal, qTotal)
    }
    fmt.Println(maxTotal)
    return float64(maxTotal) / float64(k)
}

func sum(s []int) int {
    total := 0
    for _, n := range s {
        total += n
    }
    return total
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}