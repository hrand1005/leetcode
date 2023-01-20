# """
# self is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
#class NestedInteger:
#    def isInteger(self) -> bool:
#        """
#        @return True if self NestedInteger holds a single integer, rather than a nested list.
#        """
#
#    def getInteger(self) -> int:
#        """
#        @return the single integer that self NestedInteger holds, if it holds a single integer
#        Return None if self NestedInteger holds a nested list
#        """
#
#    def getList(self) -> [NestedInteger]:
#        """
#        @return the nested list that self NestedInteger holds, if it holds a nested list
#        Return None if self NestedInteger holds a single integer
#        """

"""
class NestedIterator:
    def __init__(self, nestedList: [NestedInteger]):
        self.res = self._init_recursive(nestedList)
        self.idx = 0
    
    def _init_recursive(self, nestedList: [NestedInteger]) -> List[int]:
        res = []
        for n in nestedList:
            if n.isInteger():
                res.append(n.getInteger())
            else:
                res += self._init_recursive(n.getList())
        return res        
    
    def next(self) -> int:
        ret = self.res[self.idx]
        self.idx += 1
        return ret
    
    def hasNext(self) -> bool:
        if self.idx <= len(self.res)-1:
            return True
        return False
"""         

class NestedIterator:
    def __init__(self, nestedList: [NestedInteger]):
        self.stack = nestedList
        self.head = None

    def next(self) -> int:
        return self.head
        
    def hasNext(self) -> bool:
        while self.stack:
            n = self.stack.pop(0)
            if n.isInteger():
                self.head = n.getInteger()
                return True
            self.stack = n.getList() + self.stack
        return False    
        
# Your NestedIterator object will be instantiated and called as such:
# i, v = NestedIterator(nestedList), []
# while i.hasNext(): v.append(i.next())