OPERATORS = ["+", "-", "*", "/"]

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        for i in range(len(tokens)):
            if tokens[i] not in OPERATORS:
                tokens[i] = int(tokens[i])
                
        while len(tokens) > 1:
            index = self.find_next_operator(tokens)
            result = self.simplify(tokens[index-2], tokens[index-1], tokens[index])
            tokens = tokens[:index-2] + [result] + tokens[index+1:]
            
        return tokens[0]
        
    def find_next_operator(self, tokens: List[str]) -> int:    
        for i in range(len(tokens)):
            if tokens[i] in OPERATORS:
                return i
        return -1   
    
    def simplify(self, x: int, y: int, op: str) -> str:
        if op == "+":
            return x + y
        if op == "-":
            return x - y
        if op == "*":
            return x * y
        if op == "/":
            return int(float(x)/y)