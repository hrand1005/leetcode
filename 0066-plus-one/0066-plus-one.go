func plusOne(digits []int) []int {
    for i := len(digits)-1; i > -1; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    
    return append([]int{1}, digits...)
}