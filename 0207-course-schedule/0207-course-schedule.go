/*
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
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
    // prereqFor maps each course to a slice of
    // courses for which it is a prerequisite
    prereqFor := make(map[int][]int, numCourses)
    
    // prereqCount maps a course to its 
    // number of prerequisites
    prereqCount := make(map[int]int, numCourses)
    
    // initialize the declared graph and map
    for _, prereq := range prerequisites {
        course, req := prereq[0], prereq[1]
        prereqFor[req] = append(prereqFor[req], course)
        prereqCount[course]++
    }
    
    // initialize a queue with valid starting courses,
    // i.e. courses with no prerequisites
    queue := make([]int, 0, numCourses)
    for course, _ := range prereqFor {
        if prereqCount[course] == 0 {
            queue = append(queue, course)
        }
    }
    
    // bfs through courses, appending courses only 
    // after their prerequisites have been visited
    for len(queue) != 0 {
        req := queue[0]
        queue = queue[1:]
        for _, course := range prereqFor[req] {
            prereqCount[course]--
            if prereqCount[course] == 0 {
                queue = append(queue, course)
            }
        }
    }
    
    // check that each course had its 
    // prerequisites met during bfs traversal
    for _, count := range prereqCount {
        if count != 0 {
            return false
        }
    }
    return true
}