# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
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