class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        while len(tokens) > 1:
            index = self.find_next_operator(tokens)
            result = self.simplify(tokens[index-2], tokens[index-1], tokens[index])
            tokens = tokens[:index-2] + [result] + tokens[index+1:]
        return int(tokens[0])    
        
    def find_next_operator(self, tokens: List[str]) -> int:    
        operators = ["+", "-", "*", "/"]
        for i in range(len(tokens)):
            if tokens[i] in operators:
                return i
        return -1   
    
    def simplify(self, x: str, y: str, op: str) -> str:
        x_int, y_int = int(x), int(y)
        if op == "+":
            return str(x_int + y_int)
        if op == "-":
            return str(x_int - y_int)
        if op == "*":
            return str(x_int * y_int)
        if op == "/":
            if x_int // y_int >= 0:
                return str(x_int // y_int)
            else:
                return str(-(abs(x_int)//abs(y_int)))