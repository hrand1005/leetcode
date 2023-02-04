# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
"""
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        dummy = ListNode()
        cur = dummy
        
        while True:
            idx = self.find_min_idx(lists)
            if idx == -1:
                break
            cur.next = lists[idx]
            lists[idx] = lists[idx].next
            cur = cur.next
        
        cur.next = None
        return dummy.next
    
    def find_min_idx(self, lists: List[Optional[ListNode]]) -> int:
        min_val = 2 ** 32
        idx = 0
        all_none = True
        for i in range(len(lists)):
            if lists[i] != None and lists[i].val < min_val:
                all_none = False
                min_val = lists[i].val
                idx = i
        if all_none:
            return -1
        return idx
"""

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        if len(lists) == 0:
            return None
        if len(lists) == 1:
            return lists[0]
        
        midpoint = len(lists) // 2
        first = self.mergeKLists(lists[:midpoint])
        second = self.mergeKLists(lists[midpoint:])
        return self.mergeTwo(first, second)    
    
    def mergeTwo(self, first, second) -> Optional[ListNode]:
        dummy = ListNode()
        cur = dummy
        while first != None and second != None:
            if first.val < second.val:
                cur.next = first
                first = first.next
            else:
                cur.next = second
                second = second.next
            cur = cur.next    

        if first == None:
            cur.next = second
        else:
            cur.next = first
        
        return dummy.next