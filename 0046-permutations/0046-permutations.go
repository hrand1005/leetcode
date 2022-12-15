/*
type permTuple struct {
    perm []int
    remaining []int
}

func permute(nums []int) [][]int {
    stack := []permTuple{
        {
            perm: []int{},
            remaining: nums,
        },
    }
    
    allPerms := [][]int{}
    
    for len(stack) > 0 {
        popped := stack[0]
        stack = stack[1:]
        
        if len(popped.remaining) == 0 {
            allPerms = append(allPerms, popped.perm)
        }
        
        for i := 0; i < len(popped.remaining); i++ {
            newPerm := []int{}
            newPerm = append(newPerm, popped.perm...)
            newPerm = append(newPerm, popped.remaining[i])
            
            newRemaining := []int{}
            newRemaining = append(newRemaining, popped.remaining[:i]...)
            newRemaining = append(newRemaining, popped.remaining[i+1:]...)
            
            newPermTuple := permTuple{
                perm: newPerm,
                remaining: newRemaining,
            }
            
            stack = append(stack, newPermTuple)
        }
    }
    
    return allPerms
}
*/

func permute(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{
            { nums[0] },
        }
    }
    
    allPerms := [][]int{}
    for i := 0; i < len(nums); i++ {
        this := []int{ nums[i] }
        
        subset := []int{}
        subset = append(subset, nums[:i]...)
        subset = append(subset, nums[i+1:]...)
        
        subPerms := permute(subset)
        for _, v := range subPerms {
            newPerms := append(this, v...)
            allPerms = append(allPerms, newPerms)
        }
    }
    
    return allPerms
}