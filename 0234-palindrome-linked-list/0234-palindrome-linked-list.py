# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
"""
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        candidate = []
        cur = head
        
        while cur != None:
            candidate.append(cur.val)
            cur = cur.next
            
        for i in range(len(candidate)//2):
            if candidate[i] != candidate[len(candidate)-1-i]:
                return False
            
        return True    
"""

class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        prev = None
        slow = head
        fast = slow
        
        while fast and fast.next:
            fast = fast.next.next
            
            slow_next = slow.next
            slow.next = prev
            prev = slow
            slow = slow_next
        
        if fast:
            slow = slow.next
            
        first_half = prev
        second_half = slow
        
        while first_half and second_half:
            if first_half.val != second_half.val:
                return False
            first_half = first_half.next
            second_half = second_half.next
            
        return True    
        
            