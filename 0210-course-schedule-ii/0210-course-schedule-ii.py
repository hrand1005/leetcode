class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        if len(prerequisites) == 0 and numCourses == 1:
            return [0]
        
        prereq_for = {}
        prereq_count = {}
        
        for p in prerequisites:
            course, req = p[0], p[1]
            prereq_for[req] = prereq_for.get(req, []) + [course]
            prereq_count[course] = prereq_count.get(course, 0) + 1
        
        queue = []
        for i in range(numCourses):
            if prereq_count.get(i, 0) == 0:
                queue.append(i)
        
        class_order = []
        while queue:
            prereq = queue.pop()
            class_order.append(prereq)
            for course in prereq_for.get(prereq, []):
                prereq_count[course] -= 1
                if prereq_count[course] == 0:
                    queue.append(course)
        
        for count in prereq_count.values():
            if count != 0:
                return []
        
        return class_order