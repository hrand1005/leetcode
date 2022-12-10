/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    cur := head
    count := 1
    
    for cur.Next != nil {
        count++
        cur = cur.Next
    }
    
    if count - n == 0 {
        return head.Next
    }
    
    stitchIndex := count - n
    cur = head
    for i := 1; i < count + 1; i++ {
        if i == stitchIndex {
            tmp := cur.Next
            next := tmp.Next
            cur.Next = next
            return head
        }
        
        cur = cur.Next
    }
    
    return nil
    
}