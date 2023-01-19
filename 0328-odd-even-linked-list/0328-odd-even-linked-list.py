# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return head
        
        even = head
        odd_start = head.next
        odd = odd_start
        
        index = 2
        cur = head.next.next
        while even != None and odd != None:
            if even.next != None:
                even.next = even.next.next
                even = even.next
            if odd.next != None:    
                odd.next = odd.next.next
                odd = odd.next
        
        cur = head
        while cur.next != None:
            cur = cur.next
        
        cur.next = odd_start
        return head
