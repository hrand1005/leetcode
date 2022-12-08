class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        answer = []
        
        for i in range(1, n+1):
            v = ""

            if i % 3 == 0:
                v += "Fizz"
            if i % 5 == 0:
                v += "Buzz"    
            if len(v) == 0:
                v += str(i)

            answer.append(v)
        
        return answer