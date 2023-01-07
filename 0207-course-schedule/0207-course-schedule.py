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
        
"""

class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # represent prerequisites as a graph dict k, v 
        # where k is a vertex and v is a list of edges
        graph = {}
        # vertex to incoming edges
        requires = {}
        for p in prerequisites:
            graph[p[1]] = graph.get(p[1], []) + [p[0]]
            requires[p[0]] = requires.get(p[0], 0) + 1
        
        queue = []
        for vertex in graph.keys():
            print(vertex)
            if requires.get(vertex, 0) == 0:
                queue.append(vertex)
        
        while queue:
            cur = queue.pop()
            for e in graph.get(cur, []):
                requires[e] -= 1
                if requires[e] == 0:
                    queue.append(e)
        
        for edges in requires.values():
            if edges != 0:
                return False
        
        return True
        
        