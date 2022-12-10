/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
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
            cur.Next = cur.Next.Next
            return head
        }
        
        cur = cur.Next
    }
    
    return nil
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fast, slow := head, head
    
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    
    if fast == nil {
        return head.Next
    }
    
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    
    slow.Next = slow.Next.Next
    
    return head
}
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nodeMap := map[int]*ListNode{}
    
    cur := head
    count := 0
    for cur != nil {
        count++
        nodeMap[count] = cur
        cur = cur.Next
    }
    
    if count - n == 0 {
        return head.Next
    }
    
    stitchStart := nodeMap[count-n]
    stitchStart.Next = stitchStart.Next.Next 
    
    return head
}