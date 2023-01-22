import (
    "math/rand"
)

type Solution struct {
    original []int
}


func Constructor(nums []int) Solution {
    return Solution{
        original: nums,
    }
}


func (this *Solution) Reset() []int {
    return this.original
}


func (this *Solution) Shuffle() []int {
    pool := make([]int, len(this.original))
    copy(pool, this.original)
    shuffled := make([]int, 0, len(this.original))
    for len(pool) != 0 {
        idx := rand.Intn(len(pool))
        shuffled = append(shuffled, pool[idx])
        pool = append(pool[:idx], pool[idx+1:]...)
    }
    return shuffled
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */