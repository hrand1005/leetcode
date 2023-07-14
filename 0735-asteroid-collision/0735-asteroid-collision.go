func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0, len(asteroids))
    for _, a := range asteroids {
        stack = simulateAsteroid(stack, a)
    }
    return stack
}

func simulateAsteroid(stack []int, asteroid int) []int {
    if len(stack) == 0 {
        return []int{asteroid}
    }
    stack = append(stack, asteroid)
    for len(stack) > 1 && stack[len(stack)-2] > 0 &&  stack[len(stack)-1] < 0 {
        l, r := stack[len(stack)-2], stack[len(stack)-1]
        if l + r == 0 {
            return stack[:len(stack)-2]
        }
        survivor := l
        if abs(r) > abs(l) {
            survivor = r
        }
        stack = append(stack[:len(stack)-2], survivor)
    }
    return stack
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}