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
    next := head.Next.Next
    p1, p2 := head, head.Next
    p2.Next = p1
    p1.Next = swapPairs(next)
    return p2
}