"""
def shortSolution(x: int) -> int:
    if x == 0 or x == 1:
        return x
            
    low = 0   
    high = x // 2
    while low <= high:
        mid = (low + high) // 2
        res = mid * mid
        
        if res == x:
            return mid
        
        if res < x:
            low = mid + 1
        else:
            high = mid - 1
     
    return high
"""            
def recursiveSolution(x: int, low: int, high: int) -> int:
    mid = (low + high) // 2
    res = mid * mid
    
    # two base cases! 
    
    if res == x:
        return mid

    if res < x and (mid+1)*(mid+1) > x:
        return mid
    
    if res > x:
        return recursiveSolution(x, low, mid-1)
    else:
        return recursiveSolution(x, mid+1, high)
    
    
class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0 or x == 1:
            return x
        return recursiveSolution(x, 0, x//2)
