class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache = {}
        self.queue = []

    def get(self, key: int) -> int:
        if self.cache.get(key) != None:
            idx = self.queue.index(key)
            self.queue.pop(idx)
            self.queue.append(key)
        
        return self.cache.get(key, -1)
       
    def put(self, key: int, value: int) -> None:
        if self.cache.get(key) != None:
            idx = self.queue.index(key)
            self.queue.pop(idx)
        elif len(self.cache) == self.capacity:
            rem = self.queue.pop(0)
            self.cache.pop(rem)
            
        self.queue.append(key)
        self.cache[key] = value

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)