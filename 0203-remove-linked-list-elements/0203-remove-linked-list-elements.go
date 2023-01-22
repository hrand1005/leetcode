/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{
        Next: head,
    }
    cur := dummy
    for cur != nil && cur.Next != nil {
        if cur.Next.Val == val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return dummy.Next
}
*/

func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }
    head.Next = removeElements(head.Next, val)
    if head.Val == val {
        return head.Next
    }
    return head
}