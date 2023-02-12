# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def hasCycle(self, head):
        visited = {}
        while head != None:
            if visited.get(head, False):
                return True
            visited[head] = True
            head = head.next
        return False    
        