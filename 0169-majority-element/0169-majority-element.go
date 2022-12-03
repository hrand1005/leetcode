/*
func majorityElement(nums []int) int {
    maj := (len(nums) / 2) + 1
    occ := map[int]int{}
    
    for _, n := range nums {
        occ[n]++
        if occ[n] == maj {
            return n
        } 
    }
    return -1 // should never occur
}
*/

func majorityElement(nums []int) int {
    num, count := nums[0], 0
    
    for _, n := range nums {
        if n == num {
            count++
        } else if count > 1 {
            count--
        } else {
            num, count = n, 1
        }
    }
    
    return num
}