/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    
    return prev
}
*/

func reverse(prev, cur *ListNode) *ListNode {
    if cur == nil {
        return prev
    }
    
    next := cur.Next
    cur.Next = prev
    
    return reverse(cur, next)
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    return reverse(nil, head)
}