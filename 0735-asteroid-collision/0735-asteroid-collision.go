func asteroidCollision(asteroids []int) []int {
    var stack []int
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
    for len(stack) > 1 && willCollide(stack[len(stack)-2], stack[len(stack)-1]) {
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

func willCollide(l, r int) bool {
    return l > 0 && r < 0
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}