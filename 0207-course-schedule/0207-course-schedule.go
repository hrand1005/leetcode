func canFinish(numCourses int, prerequisites [][]int) bool {
    if len(prerequisites) == 0 {
        return true
    }
    
    finishable := make(map[int]bool, numCourses)
    for i := 0; i < numCourses; i++ {
        visited := make(map[int]bool, numCourses)
        if !canFinishRecursive(i, visited, finishable, prerequisites) {
            return false
        }
    }
    
    return true
}


func canFinishRecursive(course int, visited, finishable map[int]bool, prerequisites [][]int) bool {
    if finishable[course] {
        return true
    }
    
    if visited[course] {
        return false
    }
    
    visited[course] = true
    for _, p := range prerequisites {
        if p[0] == course {
            if !canFinishRecursive(p[1], visited, finishable, prerequisites) {
                return false
            }
        }
    }
    
    finishable[course] = true
    return true
}
