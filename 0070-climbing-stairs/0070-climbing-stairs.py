def recursiveSolution(i: int, one_before: int, two_before: int) -> int:
    if i == 0:
        return one_before + two_before
    return recursiveSolution(i-1, two_before + one_before, one_before)

# way too slow
def recursiveSolution2(n: int) -> int:
    if n < 3:
        return n
    
    return recursiveSolution2(n-1) + recursiveSolution2(n-2)

def forLoopSolution(n: int) -> int:
    two_before = 1
    one_before = 2
    for i in range(3, n+1):
        cur = two_before + one_before
        two_before = one_before
        one_before = cur

    return cur

def mapSolution(n: int) -> int:
    distinct_ways = {
        1: 1,
        2: 2,
    }
    
    for i in range(3, n+1):
        distinct_ways[i] = distinct_ways[i-1] + distinct_ways[i-2]
        
    return distinct_ways[n]    

class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 3:
            return n
        
        return recursiveSolution(n-3, 2, 1)
        # return recursiveSolution2(n)
        # return forLoopSolution(n)
        # return mapSolution(n)
        
