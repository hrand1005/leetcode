"""
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # can finish current course's prerequisites?
        self.finished = set()
        
        for i in range(numCourses):
            if not self.canFinishRecursive(i, prerequisites):
                return False
        return True   
    
    def canFinishRecursive(self, course: int, prerequisites: List[List[int]]) -> bool:
        if course in self.finished:
            return True
        
        for p in prerequisites:
            if p[0] == course:
                if not self.canFinishRecursive(p[1], prerequisites):
                    return False
        
        self.finished.add(course)       
        return True
"""
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        self.finished = set()
        
        for i in range(numCourses):
            if self.contains_cycle(i, set(), prerequisites):
                return False
        
        return True
    
    def contains_cycle(self, course: int, visited: Set[int], prerequisites: List[List[int]]) -> bool:
        if course in self.finished:
            return False
        
        if course in visited:
            return True
        
        visited.add(course)
        
        for p in prerequisites:    
            if course == p[0]:
                if self.contains_cycle(p[1], visited, prerequisites):
                    return True
                    
        self.finished.add(course)            
        return False            