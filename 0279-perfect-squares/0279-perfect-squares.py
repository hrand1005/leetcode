"""
class Solution:
    def numSquares(self, n: int) -> int:
        num_squares = 0
        squares = []
        while 0 < n:
            tmp = n
            while n < tmp * tmp:
                tmp -= 1
            n -= tmp * tmp
            num_squares += 1
        return num_squares
"""   

class Solution:
    def numSquares(self, n: int) -> int:
        self.cache = {}
        squares = []
        for i in range(1, n):
            if i * i < n:
                squares.append(i * i)
                self.cache[i*i] = 1
            if i * i == n:
                return 1
        return self.num_squares_recursive(n, squares)
        
    def num_squares_recursive(self, n: int, squares: List[int]) -> int:
        if self.cache.get(n) != None:
            return self.cache[n]
        
        # all 1's, for example
        min_num = n
        for s in squares:
            if n < s:
                break
            if s == n:
                return 1
            min_num = min(min_num, 1 + self.num_squares_recursive(n-s, squares))
        
        self.cache[n] = min_num
        return min_num

"""
class Solution:
    def numSquares(self, n: int) -> int:
        table = [0] * n + 1
        table[1] = 1
        for i in range(n):
            if i*i <= n:
                table[i*i] = 1
            table[i] = 
"""