"""
class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 3:
            return 0
        if n == 3:
            return 1
        
        primes = 1
        for i in range(3, n):
            prime = True
            for j in range(2, i):
                if i % j == 0:
                    prime = False
                    break
            if prime:
                primes += 1
        
        return primes
"""

class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 3:
            return 0
        
        table = [True] * n
        table[0], table[1] = False, False
        
        for i in range(2, n):
            if table[i]:
                cur = i * i
                while cur < n:
                    table[cur] = False
                    cur += i
        
        return sum(table)
                    