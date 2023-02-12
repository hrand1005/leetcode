# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

"""
class Solution(object):
    def getIntersectionNode(self, headA, headB):
        visited = {}
        while headA != None:
            visited[headA] = True
            headA = headA.next
        
        while headB != None:
            if visited.get(headB, False):
                return headB
            headB = headB.next
        return None    
"""

# O(1) Space
class Solution(object):
    def getIntersectionNode(self, headA, headB):
        a_start = headA
        b_start = headB
        
        while headA != headB:
            if headA is None:
                headA = b_start
                continue
            if headB is None:
                headB = a_start
                continue
            headA = headA.next
            headB = headB.next
        
        return headA