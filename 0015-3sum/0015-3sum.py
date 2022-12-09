from collections import Counter
# Brute force solution 8 min
"""
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        solutions = []
        
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                for k in range(j+1, len(nums)):
                    if nums[i] + nums[j] + nums[k] == 0:
                        duplicate = False
                        for s in solutions:
                            if Counter([nums[i], nums[j], nums[k]]) == Counter(s):
                                duplicate = True
                                
                        if not duplicate:
                            solutions.append([nums[i], nums[j], nums[k]])
        
        return solutions
"""

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        negative, zero, positive = [], [], []
        
        for n in nums:
            if n < 0:
                negative.append(n)
            elif n == 0:
                zero.append(n)
            else:
                positive.append(n)
        
        n_set, p_set = set(negative), set(positive)
        solutions = set()
            
        if len(zero) > 0:
            for n in n_set:
                comp = n*-1
                if comp in p_set:
                    solutions.add((n, 0, n*-1))
            if len(zero) > 2:
                solutions.add((0, 0, 0))
                    
        for i in range(len(negative)):
            for j in range(i+1, len(negative)):
                comp = (negative[i] + negative[j]) * -1
                if comp in p_set:
                    solutions.add(tuple(sorted([negative[i], negative[j], comp])))

        for i in range(len(positive)):
            for j in range(i+1, len(positive)):
                comp = (positive[i] + positive[j]) * -1
                if comp in n_set:
                    solutions.add(tuple(sorted([positive[i], positive[j], (positive[i] + positive[j]) * -1])))
        
        return solutions