type key struct {
    i, j int
}

type result struct {
    a, b, c, d int
}

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    
    subproblems := make(map[key]int)
    for i := 0; i < len(nums); i++ {
        for j := i+3; j < len(nums); j++ {
            subproblems[key{i, j}] = target - (nums[i] + nums[j])
        }
    }
    
    resultSet := make(map[result]bool)
    for k, v := range subproblems {
        start, end := k.i+1, k.j
        middleSet := twoSum(nums[start:end], v)
        for _, s := range middleSet {
            resultSet[result{nums[k.i], s[0], s[1], nums[k.j]}] = true
        }
    }
    
    quads := make([][]int, 0)
    for res, _ := range resultSet {
        quads = append(quads, []int{res.a, res.b, res.c, res.d})
    }
    
    return quads
}

func twoSum(nums []int, target int) [][]int {
    pairs := make([][]int, 0)
    occ := make(map[int]bool)
    for _, n := range nums {
        if occ[target-n] {
            pairs = append(pairs, []int{target-n, n})
        }
        occ[n] = true
    }
    return pairs
}
