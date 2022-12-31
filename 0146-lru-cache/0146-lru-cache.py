"""
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

"""
class Node:
    def __init__(self, key=-1, val=-1, next=None, prev=None):
        self.val = val
        self.key = key
        self.next = next
        self.prev = prev

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.map = {}
        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head

    def get(self, key: int) -> int:
        node = self.map.get(key)
        if node != None:
            self._remove_node(node)
            self._add_to_tail(node)
            return node.val
        return -1
        
    def put(self, key: int, value: int) -> None:
        node = self.map.get(key)
        if node != None:
            self._remove_node(node)
            node.val = value
        else:
            if len(self.map) == self.capacity:
                self.map.pop(self.head.next.key)
                self._remove_node(self.head.next)
            node = Node(key=key, val=value)
        self._add_to_tail(node)
        self.map[key] = node
            
    def _remove_node(self, node: Node):
        prev = node.prev
        next = node.next
        prev.next = next
        next.prev = prev
   
    def _add_to_tail(self, node: Node):
        prev = self.tail.prev
        prev.next = node
        node.prev = prev
        node.next = self.tail
        self.tail.prev = node

