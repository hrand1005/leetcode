/*
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

func combinationSum(candidates []int, target int) [][]int {
    combinations := make([][]int, 0)
    for i, c := range candidates {
        if c == target {
            combinations = append(combinations, []int{c})
        }
        if c < target {
            subcombos := combinationSum(candidates[i:], target-c)
            for _, s := range subcombos {
                full := append(s, c)
                combinations = append(combinations, full)
            }
        } 
    }
    return combinations
}
*/

func combinationSum(candidates []int, target int) [][]int {
    combos := make([][]int, 0)
    
    var dfs func([]int, int, []int) 
    dfs = func(cands []int, targ int, path []int) {
        if targ < 0 {
            return
        }
        if targ == 0 {
            combos = append(combos, path)
            return
        }
        for i, c := range cands {
            newPath := make([]int, len(path))
            copy(newPath, path)
            dfs(cands[i:], targ-c, append(newPath, c))
        }
        return
    }
    
    dfs(candidates, target, []int{})
    return combos
}