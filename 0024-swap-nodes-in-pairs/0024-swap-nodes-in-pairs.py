# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        
        new_head = head.next
        cur = head
        prev = None
        while cur is not None and cur.next is not None:
            p1, p2 = cur, cur.next
            cur = cur.next.next
            p2.next = p1
            p1.next = cur
            if prev:
                prev.next = p2
            prev = p1    
            
        return new_head   