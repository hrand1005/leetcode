class Solution:
    """
    def plusOne(self, digits: List[int]) -> List[int]:
        for i in range(len(digits)-1, -1, -1):
            if digits[i] == 9:
                digits[i] = 0
            else:
                digits[i] += 1
                return digits
        
        if digits[0] == 0:
            digits.insert(0, 1)
            
        return digits   
    """    
    def plusOne(self, digits: List[int]) -> List[int]:
        if digits[-1] == 9:
            if len(digits) == 1:
                return [1, 0]
            
            return self.plusOne(digits[:-1]) + [0]
        
        digits[-1] += 1
        return digits
        
             