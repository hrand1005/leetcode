func findOrder(numCourses int, prerequisites [][]int) []int {
    postreqs := make(map[int][]int, numCourses)
    prereqCount := make(map[int]int, numCourses)
    for _, p := range prerequisites {
        post, pre := p[0], p[1]
        postreqs[pre] = append(postreqs[pre], post)
        prereqCount[post]++
    }
    
    queue := make([]int, 0, numCourses)
    for i := 0; i < numCourses; i++ {
        if prereqCount[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    courseOrder := make([]int, 0, numCourses)
    for len(queue) != 0 {
        course := queue[0]
        queue = queue[1:]
        courseOrder = append(courseOrder, course)
        for _, post := range postreqs[course] {
            prereqCount[post]--
            if prereqCount[post] == 0 {
                queue = append(queue, post)
            }
        }
    }
    
    for _, count := range prereqCount {
        if count != 0 {
            return nil
        }
    }
    
    return courseOrder
}