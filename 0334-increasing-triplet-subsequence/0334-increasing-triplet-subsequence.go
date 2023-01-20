import (
    "math"
)

func increasingTriplet(nums []int) bool {
    low, mid := math.MaxInt32, math.MaxInt32
    for _, n := range nums {
        if n <= low {
            low = n
        } else if n <= mid {
            mid = n
        } else {
            return true
        }
    }
    return false
}