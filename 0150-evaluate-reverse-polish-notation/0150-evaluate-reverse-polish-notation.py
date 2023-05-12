class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        for t in tokens:
            if t == "+":
                a, b = stack[-1], stack[-2]
                stack = stack[:-2] + [a+b]
            elif t == "-":
                a, b = stack[-1], stack[-2]
                stack = stack[:-2] + [b-a]
            elif t == "*":
                a, b = stack[-1], stack[-2]
                stack = stack[:-2] + [a*b]
            elif t == "/":
                a, b = stack[-1], stack[-2]
                stack = stack[:-2] + [int(b/a)]
            else:
                stack.append(int(t))
        return stack[-1]