# TWO QUEUES
"""
class MyStack:

    def __init__(self):
        self.q1 = []
        self.q2 = []
        self.last = None
        
    def push(self, x: int) -> None:
        self.q1.append(x)
        self.last = x

    def pop(self) -> int:
        for i in range(len(self.q1)-1):
            self.last = self.q1.pop(0)
            self.q2.append(self.last)
        ret = self.q1.pop(0)    
        self.q1, self.q2 = self.q2, self.q1
        return ret

    def top(self) -> int:
        return self.last

    def empty(self) -> bool:
        return len(self.q1) == 0
"""

# ONE QUEUE
class MyStack:

    def __init__(self):
        self.queue = []
        self.last = None
        
    def push(self, x: int) -> None:
        self.queue.append(x)
        self.last = x

    def pop(self) -> int:
        for _ in range(len(self.queue)-1):
            self.last = self.queue.pop(0)
            self.queue.append(self.last)
        return self.queue.pop(0)    

    def top(self) -> int:
        return self.last

    def empty(self) -> bool:
        return len(self.queue) == 0

# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()