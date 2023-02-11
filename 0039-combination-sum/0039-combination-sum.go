type withTarget struct {
    remaining int
    combo []int
    pool []int
}

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    
    combinations := make([][]int, 0)
    queue := []*withTarget{
        {
            remaining: target,
            pool: candidates,
        },
    }
    for len(queue) != 0 {
        this := queue[0]
        queue = queue[1:]
        
        for i, c := range this.pool {
            if c < this.remaining {
                newCombo := make([]int, len(this.combo))
                copy(newCombo, this.combo)
                newPool := make([]int, len(this.pool))
                copy(newPool, this.pool)
                queue = append(queue, &withTarget{
                    remaining: this.remaining - c,
                    combo: append(newCombo, c),
                    pool: newPool[i:],
                })
            }
            if c == this.remaining {
                combinations = append(combinations, append([]int{c}, this.combo...))
            }
        }
    }
    
    return combinations
}
