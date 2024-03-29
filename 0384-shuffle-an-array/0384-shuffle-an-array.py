class Solution:

    def __init__(self, nums: List[int]):
        self.original = nums

    def reset(self) -> List[int]:
        return self.original
        

    def shuffle(self) -> List[int]:
        shuffled = []
        pool = self.original.copy()
        while pool:
            n = pool.pop(random.randrange(0, len(pool)))
            shuffled.append(n)
        return shuffled    
        
        


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.reset()
# param_2 = obj.shuffle()