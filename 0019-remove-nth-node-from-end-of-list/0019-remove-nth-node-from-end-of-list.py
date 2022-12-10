# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
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
                
                
            
            
        
        
        