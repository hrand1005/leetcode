/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    slow := head
    fast := slow.Next
    
    for fast != nil {
        if fast == slow { 
            return true
        }
        
        fast = fast.Next 
        if fast == nil {
            return false
        }
        
        if fast == slow { 
            return true
        }
        
        slow = slow.Next 
        fast = fast.Next
    }
    
    return false
}