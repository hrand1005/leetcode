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
