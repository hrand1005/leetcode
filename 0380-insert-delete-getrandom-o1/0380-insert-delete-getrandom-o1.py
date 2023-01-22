"""
class RandomizedSet:

    def __init__(self):
        self.set = {}

    def insert(self, val: int) -> bool:
        not_present = False
        if self.set.get(val) == None:
            not_present = True
            self.set[val] = True
        return not_present
        

    def remove(self, val: int) -> bool:
        if self.set.get(val) == None:
            return False
        self.set.pop(val)
        return True
        
    def getRandom(self) -> int:
        return list(self.set.keys())[random.randint(0, len(self.set)-1)]
"""       

class RandomizedSet:

    def __init__(self):
        self.map = {}
        self.lst = []

    def insert(self, val: int) -> bool:
        if self.map.get(val) == None:
            self.map[val] = len(self.lst)
            self.lst.append(val)
            return True
        return False
        
    def remove(self, val: int) -> bool:
        if self.map.get(val) == None:
            return False
        
        idx = self.map[val]
        self.lst[idx] = self.lst[-1]
        self.map[self.lst[idx]] = idx
        
        self.lst.pop()
        self.map.pop(val)
        
        return True
        
    def getRandom(self) -> int:
        return self.lst[random.randint(0, len(self.lst)-1)]

# Your RandomizedSet object will be instantiated and called as such:
# obj = RandomizedSet()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()