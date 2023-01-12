/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
func deleteNode(node *ListNode) {
    current := node
    for current.Next.Next != nil {
        current.Val = current.Next.Val
        current = current.Next
    }
    current.Val = current.Next.Val
    current.Next = nil
}

func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}
*/

func deleteNode(node *ListNode) {
    *node = *node.Next
}