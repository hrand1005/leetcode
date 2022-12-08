/*
func missingNumber(nums []int) int {
    occurred := map[int]int{}

    for _, v := range nums {
        occurred[v]++
    }

    for i := 0; i < len(nums)+1; i++ {
        if _, ok := occurred[i]; !ok {
            return i
        }
    }

    return -1
}
*/

func missingNumber(nums []int) int {
    got := 0
    exp := len(nums)
    for i := 0; i < len(nums); i++ {
        got += nums[i]
        exp += i
    }
    return exp - got
}