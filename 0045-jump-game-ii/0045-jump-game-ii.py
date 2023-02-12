MAX_INT = 2 ** 32

class Solution:
    def jump(self, nums: List[int]) -> int:
        self.cache = {}
        return self.jump_recursive(nums, 0)
    
    def jump_recursive(self, nums: List[int], pos: int) -> int:
        if self.cache.get(pos) != None:
            return self.cache[pos]
        
        if pos == len(nums)-1:
            return 0
        
        scope = nums[pos]
        if scope+pos >= len(nums)-1:
            return 1
        
        min_jumps = MAX_INT
        for i in range(pos+1, pos+scope+1):
            min_jumps = min(min_jumps, self.jump_recursive(nums, i))
        
        min_jumps += 1
        self.cache[pos] = min_jumps
        return min_jumps 