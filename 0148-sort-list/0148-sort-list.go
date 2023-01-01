/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    left, right := splitList(head)
    sortedLeft := sortList(left)
    sortedRight := sortList(right)
    return merge(sortedLeft, sortedRight)
}

func splitList(head *ListNode) (*ListNode, *ListNode) {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    right := slow.Next
    slow.Next = nil
    return head, right
}

func merge(left, right *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for left != nil && right != nil {
        if left.Val < right.Val {
            cur.Next = left
            left = left.Next
        } else {
            cur.Next = right
            right = right.Next
        }
        cur = cur.Next
    }
    
    if left != nil {
        cur.Next = left
    } else {
        cur.Next = right
    }
    
    return dummy.Next
}