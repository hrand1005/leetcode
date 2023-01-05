/*
func trailingZeroes(n int) int {
    zeroes := 0
    for 0 < n {
        n /= 5
        zeroes += n
    }
    return zeroes
}
*/

func trailingZeroes(n int) int {
    divisor := 5
    zeroes := 0
    for divisor <= n {
        zeroes += n / divisor
        divisor *= 5
    }
    return zeroes
}