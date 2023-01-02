class MinStack:

    def __init__(self):
        self.min_val = 0
        self.stack = []

    def push(self, val: int) -> None:
        if len(self.stack) == 0 or val < self.min_val:
            self.min_val = val
        self.stack.append(val)    

    def pop(self) -> None:
        val = self.stack.pop()
        if val == self.min_val:
            self._set_next_min()

    def top(self) -> int:
        return self.stack[-1]
        
    def getMin(self) -> int:
        return self.min_val
    
    def _set_next_min(self) -> int:
        if len(self.stack) != 0:
            self.min_val = self.stack[0]
            for n in self.stack:
                self.min_val = min(self.min_val, n)

# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()