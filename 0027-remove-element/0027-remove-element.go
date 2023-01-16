/*
func removeElement(nums []int, val int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        if nums[left] == val {
            nums[left], nums[right] = nums[right], nums[left]
            right--
        } else {
            left++
        }
    }
    return left
}
*/

func removeElement(nums []int, val int) int {
    i := 0
    for _, n := range nums {
        if n != val {
            nums[i] = n
            i++
        }
    }
    return i
}