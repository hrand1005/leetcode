# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

"""
class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        if headA == headB:
            return headA
        
        encountered_nodes = {}
        cur1 = headA
        cur2 = headB
        while cur1 != None or cur2 != None:
            if cur1:
                if encountered_nodes.get(cur1):
                    return cur1
                encountered_nodes[cur1] = 1
                cur1 = cur1.next
                
            if cur2:
                if encountered_nodes.get(cur2):
                    return cur2
                encountered_nodes[cur2] = 1
                cur2 = cur2.next
        
        return None
"""
class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        cur1 = headA
        cur2 = headB
        
        while cur1 != cur2:
            if cur1 == None:
                cur1 = headB
            else:
                cur1 = cur1.next
                
            if cur2 == None:
                cur2 = headA
            else:
                cur2 = cur2.next
                
        return cur1