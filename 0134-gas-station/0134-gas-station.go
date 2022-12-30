/*
func canCompleteCircuit(gas []int, cost []int) int {
    if sum(gas) < sum(cost) {
        return -1
    }
    
    lastDiff := -1
    for i := 0; i < len(gas); i++ {
        if cost[i] <= gas[i] && lastDiff < 0 {
            if canComplete(i, gas, cost) {
                return i
            }
        }
    }
    return -1
}

func canComplete(start int, gas, cost []int) bool {
    tank := 0
    for i := 0; i < len(gas); i++ {
        index := (start + i) % len(gas)
        tank += gas[index] - cost[index]
        if tank < 0 {
            return false
        }
    }
    return true
}

func sum(s []int) int {
    sum := 0
    for i := 0; i < len(s); i++ {
        sum += s[i]
    }
    return sum
}

func canCompleteCircuit(gas []int, cost []int) int {
    start, tank, total := 0, 0, 0
    for i := 0; i < len(gas); i++ {
        total += gas[i] - cost[i]
        tank += gas[i] - cost[i]
        if tank < 0 {
            tank = 0
            start = i + 1
        }
    }
    if total < 0 {
        return -1
    }
    return start
}
*/

func canCompleteCircuit(gas []int, cost []int) int {
    diff := make([]int, len(gas))
    for i := 0; i < len(diff); i++ {
        diff[i] = gas[i] - cost[i]
    }
    
    start, maxSub, total := 0, 0, 0
    for i := 0; i < len(diff); i++ {
        total += diff[i]
        maxSub += diff[i]
        if diff[i] > maxSub {
            start = i
            maxSub = diff[i]
        }
    }
    
    if total < 0 {
        return -1
    }
    
    return start
}
