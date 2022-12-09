func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    maxArea := 0
    
    for l < r {
        length := r - l
        depth := height[l]
        
        if height[r] < height[l] {
            depth = height[r]
        }
        
        area := length * depth
        
        if maxArea < area {
            maxArea = area
        }
        
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    
    return maxArea
}