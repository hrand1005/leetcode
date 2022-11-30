# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        val_set = {}
        cur = head
        while cur:
            if val_set.get(cur):
                return True
            val_set[cur] = 1
            cur = cur.next
        
        return False