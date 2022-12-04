func calcNext(n int) int {
    digits := []int{}
    for n >= 10 {
        digits = append(digits, n % 10)
        n /= 10
    }
    digits = append(digits, n)
    
    next := 0
    for _, v := range digits {
        next += v * v
    }
    
    return next
}

/*
func isHappy(n int) bool {
    occurred := map[int]bool{}
    
    for n != 1 {
        if _, ok := occurred[n]; ok {
            return false
        }
        occurred[n] = true
        n = calcNext(n)
    }
    
    return true
}
*/

func isHappy(n int) bool {
    slow := n
    fast := calcNext(slow)
    
    for fast != 1 {
        if fast == slow {
            return false
        }
        
        slow = calcNext(slow)
        fast = calcNext(fast)
        
        if fast == slow {
            return false
        }
        
        fast = calcNext(fast)
    }
    
    return true
}