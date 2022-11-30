/*
func singleNumber(nums []int) int {
    occ := map[int]int{}
    for _, n := range nums {
        occ[n] += 1
    }
    for k, v := range occ {
        if v == 1 {
            return k
        }
    }
    return 0
}
*/

func singleNumber(nums []int) int {
    res := 0
    for _, n := range nums {
        res ^= n
    }
    return res
}