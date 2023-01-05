func trailingZeroes(n int) int {
    zeroes := 0
    for 0 < n {
        n /= 5
        zeroes += n
    }
    return zeroes
}