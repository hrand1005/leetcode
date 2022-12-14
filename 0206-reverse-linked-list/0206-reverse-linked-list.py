# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
"""
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return head
        
        prev = head
        cur = head.next
        head.next = None
        
        while cur.next != None:
            next_node = cur.next
            cur.next = prev
            prev = cur
            cur = next_node
            
        cur.next = prev     
        
        return cur
""" 

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None:
            return head
        
        return self.reverse(None, head)
    
    def reverse(self, prev, cur):
        if cur.next == None:
            cur.next = prev
            return cur
        
        nxt = cur.next
        cur.next = prev
        return self.reverse(cur, nxt)
        
        
            
        