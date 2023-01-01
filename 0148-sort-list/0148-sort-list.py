# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
"""
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return head
        
        left_head, right_head = self.split_linked_list(head)
        left = self.sortList(left_head)
        right = self.sortList(right_head)
        return self.merge(left, right)
    
    def split_linked_list(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head
        count = 0
        while cur:
            count += 1
            cur = cur.next
        
        cur = head
        count = count // 2
        for i in range(1, count):
            cur = cur.next
        
        right = cur.next
        cur.next = None
        return head, right
    
    def merge(self, left: Optional[ListNode], right: Optional[ListNode]) -> Optional[ListNode]:
        head = None
        if left.val < right.val:
            head = left
            left = left.next
        else:
            head = right
            right = right.next
            
        cur = head
        while left != None and right != None:
            if left.val < right.val:
                cur.next = left
                left = left.next
            else:
                cur.next = right
                right = right.next
            cur = cur.next
        
        if left != None:
            cur.next = left
        elif right != None:
            cur.next = right
        
        return head
"""
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return head
        
        left_head, right_head = self.split(head)
        left = self.sortList(left_head)
        right = self.sortList(right_head)
        return self.merge(left, right)
    
    def split(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow, fast = head, head.next
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        right = slow.next
        slow.next = None
        return head, right
    
    def merge(self, left: Optional[ListNode], right: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode()
        cur = dummy
        while left != None and right != None:
            if left.val < right.val:
                cur.next = left
                left = left.next
            else:
                cur.next = right
                right = right.next
            cur = cur.next
        
        cur.next = left or right
        #if left != None:
        #    cur.next = left
        #elif right != None:
        #    cur.next = right
        
        return dummy.next