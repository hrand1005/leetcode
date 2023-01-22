# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        if head is None:
            return head
        
        cur = head
        while cur is not None and cur.val == val:
            cur = cur.next
        
        new_head = cur
        while cur is not None:
            while cur.next is not None and cur.next.val == val:
                cur.next = cur.next.next
            cur = cur.next    
        return new_head    
            
        
        