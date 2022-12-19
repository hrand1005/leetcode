"""
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        result = []
        
        i = 0
        while i < len(intervals):
            low, high = intervals.pop(0)
            
            j = 0
            changed = False
            while j < len(intervals):
                if high < intervals[j][0] or intervals[j][1] < low:
                    j += 1
                else:    
                    this_low, this_high = intervals.pop(j)
                    low = min(low, this_low)
                    high = max(high, this_high)
                    changed = True
            
            if changed: 
                intervals = [[low, high]] + intervals
            else:    
                result.append([low, high])
            
        return result
"""

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key= lambda x : x[0])
        result = intervals[:1]
        
        for n in intervals[1:]:
            if n[0] <= result[-1][1]:
                result[-1][1] = max(n[1], result[-1][1]) 
            else:
                result.append(n)
        
        return result
                