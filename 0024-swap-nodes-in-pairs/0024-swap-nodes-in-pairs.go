/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    var prev *ListNode
    cur := head
    newHead := head.Next
    
    for cur != nil && cur.Next != nil {
        p1, p2 := cur, cur.Next
        cur = cur.Next.Next
        p2.Next = p1
        p1.Next = cur
        if prev != nil {
            prev.Next = p2  
        }
        prev = p1
    }
    
    return newHead
}
    
/*    
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    next := head.Next.Next
    p1, p2 := head, head.Next
    p2.Next = p1
    p1.Next = swapPairs(next)
    return p2
}
*/