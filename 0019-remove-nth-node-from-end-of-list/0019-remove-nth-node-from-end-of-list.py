# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
"""
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        if head.next == None:
            return None
        
        cur = head
        count = 1
        
        while cur.next != None:
            count += 1
            cur = cur.next
        
        prev = head
        cur = head.next
        target = count - n + 1
        
        if target == 1:
            return head.next
        
        for i in range(2, count+1):
            if i == target:
                prev.next = cur.next
                return head
                
            prev = prev.next    
            cur = cur.next
        
        return head

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        fast, slow = head, head
        
        for _ in range(n):
            fast = fast.next
            
        if not fast:
            return head.next
        
        while fast.next:
            slow = slow.next
            fast = fast.next
        
        rem = slow.next
        next_node = rem.next
        slow.next = next_node
        
        return head
"""                

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        node_map = {}
        
        i = 0
        cur = head
        while cur:
            i += 1
            node_map[i] = cur
            cur = cur.next
        
        if (i-n) == 0:
            return head.next
        
        before = node_map.get(i-n)
        after = node_map.get(i-n+2)
        before.next = after
        
        return head
        
        
            
            
        
        
        