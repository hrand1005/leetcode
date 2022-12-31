"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if head == None:
            return None
        
        head_copy = Node(head.val)
        cur, cur_copy = head, head_copy
        orig_list, copy_list = [], []
        while cur:
            if cur.next != None:
                cur_copy.next = Node(cur.next.val)
            
            orig_list.append(cur)
            copy_list.append(cur_copy)
            cur = cur.next
            cur_copy = cur_copy.next
        
        for i in range(len(orig_list)):
            orig_rand = orig_list[i].random
            if orig_rand == None:
                copy_list[i].random = None
            else:    
                for j in range(len(orig_list)):
                    if orig_list[j] == orig_rand:
                        copy_list[i].random = copy_list[j]
        
        return head_copy
"""

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if head == None:
            return None
        
        head_copy = Node(head.val)
        cur, cur_copy = head, head_copy
        orig_to_copy = {}
        while cur:
            if cur.next != None:
                cur_copy.next = Node(cur.next.val)
            orig_to_copy[cur] = cur_copy
            cur = cur.next    
            cur_copy = cur_copy.next
        
        cur, cur_copy = head, head_copy
        while cur:
            if cur.random == None:
                cur_copy.random = None
            else:    
                cur_copy.random = orig_to_copy[cur.random]
            cur = cur.next
            cur_copy = cur_copy.next
        
        return head_copy