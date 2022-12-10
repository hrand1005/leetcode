"""
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        return self.generate(n, n, '', [])
        
    def generate(self, lefts, rights, current, answers):
        if lefts > rights:
            return []
        
        if lefts == 0 and rights == 0:
            answers.append(current)
            return answers
        
        left_combos, right_combos = [], []
        if lefts > 0:
            self.generate(lefts-1, rights, current + '(', answers)
        if rights > 0:    
            self.generate(lefts, rights-1, current + ')', answers)
        
        return answers
"""

class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        left, right = n, n
        stack = [('(', left-1, right)]
        combos = []
        
        while stack:
            p = stack.pop()
            if p[1] > p[2]:
                continue
            
            if p[1] == 0 and p[2] == 0:
                combos.append(p[0])
                
            if p[1] > 0:
                new_left = (p[0] + '(', p[1] - 1, p[2])
                stack.append(new_left)
                
            if p[2] > 0:
                new_right = (p[0] + ')', p[1], p[2]-1)
                stack.append(new_right)
            
        return combos