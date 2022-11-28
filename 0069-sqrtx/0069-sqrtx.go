func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }
    return findRecursively(x, 0, x/2)
}

func findRecursively(x, low, high int) int {
    midpoint := (high + low) / 2
    product := midpoint * midpoint

    if product == x {
        return midpoint
    }

    if product < x && x < (midpoint+1) * (midpoint+1) {
        return midpoint
    }
    
    if product > x {
        return findRecursively(x, low, midpoint-1)
    }
    
    return findRecursively(x, midpoint+1, high)
}

/*
func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }
    
    low := 0
    high := x / 2
    for low <= high {
        midpoint := (high + low) / 2
        product := midpoint * midpoint
        
        if product == x {
            return midpoint
        }
        
        if product < x && x < (midpoint+1) * (midpoint+1) {
            return midpoint
        }
        
        if product > x {
            high = midpoint - 1
        } else {
            low = midpoint + 1
        }
    }
    
    return high
}
*/