/*
func getSum(a int, b int) int {
    if b == 0 {
        return a
    }
    return getSum(a^b, (a&b)<<1)
}
*/

func getSum(a int, b int) int {
    for b != 0 {
        a, b = (a^b), (a&b)<<1
    }
    return a
}