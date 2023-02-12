# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        visited = {}
        while headA != None:
            visited[headA] = True
            headA = headA.next
        
        while headB != None:
            if visited.get(headB, False):
                return headB
            headB = headB.next
            
        return None    
        
        