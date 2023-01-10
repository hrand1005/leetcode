class Solution:
    def calculate(self, s: str) -> int:
        num = 0
        op = "+"
        stack = []
        
        for c in s+"+":
            if c.isdigit():
                num = num * 10 + int(c)
            elif c in "+-*/":
                if op == "+":
                    stack.append(num)
                if op == "-":
                    stack.append(-num)
                if op == "*":
                    prev = stack.pop()
                    stack.append(prev * num)
                if op == "/":
                    prev = stack.pop()
                    quotient = abs(prev) // num
                    if prev < 0:
                        quotient = -quotient
                    stack.append(quotient)
                op = c    
                num = 0
        
        return sum(stack)
