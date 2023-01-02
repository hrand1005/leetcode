"""
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
"""

class MinStack:

    def __init__(self):
        # stack of (min, cur) tuples
        self.stack = []

    def push(self, val: int) -> None:
        if len(self.stack) == 0:
            self.stack.append((val, val))
        else:
            prev_min = self.stack[-1][0]
            new_min = min(prev_min, val)
            self.stack.append((new_min, val))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][1]
        
    def getMin(self) -> int:
        return self.stack[-1][0]
