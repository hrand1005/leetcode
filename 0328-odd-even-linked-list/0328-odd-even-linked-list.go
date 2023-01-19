/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    odd := head
    even := head.Next
    evenStart := even
    for even != nil && even.Next != nil {
        odd.Next = odd.Next.Next
        odd = odd.Next
        even.Next = even.Next.Next
        even = even.Next
    }
    odd.Next = evenStart
    return head
}