func minMoves(nums []int) int {
    return sumOf(nums) - len(nums) * minOf(nums)
}

func sumOf(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}

func minOf(nums []int) int {
    min := nums[0]
    for _, n := range nums[1:] {
        if n < min {
            min = n
        }
    }
    return min
}