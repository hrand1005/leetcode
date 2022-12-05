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
            # increment fast at double the rate
            fast = fast.next.next
            
            # reverse each element in the first half
            slow_next = slow.next
            slow.next = prev
            prev = slow
            slow = slow_next
        
        # for odd numbers of elements, don't compare the middle element
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
        
            