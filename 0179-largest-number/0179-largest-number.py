class Solution:
    def largestNumber(self, nums: List[int]) -> str:
        nums = self.mergesort(nums, 0, len(nums)-1)
        result = ""
        for n in nums:
            result += str(n)
        return str(int(result))
    
    def mergesort(self, nums: List[int], left: int, right: int) -> List[int]:
        if right <= left:
            return [nums[right]]
        
        midpoint = (left + right) // 2
        first = self.mergesort(nums, left, midpoint)
        second = self.mergesort(nums, midpoint+1, right)
        
        return self.merge(first, second)
   
    def merge(self, first: List[int], second: List[int]) -> List[int]:
        full = []
        while len(first) != 0 and len(second) != 0:
            if self.less_than(first[0], second[0]):
                full.append(second.pop(0))
            else:
                full.append(first.pop(0))

        full += first or second
        return full
    
    def less_than(self, a: int, b: int) -> int:
        str_a, str_b = str(a), str(b)
        if int(str_a+str_b) < int(str_b+str_a):
            return True
        return False
