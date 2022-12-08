/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}

    cur := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        cur.Next = &ListNode{}
        cur = cur.Next

        digitSum := carry

        if l1 != nil {
            digitSum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            digitSum += l2.Val
            l2 = l2.Next
        }

        carry = digitSum / 10
        digitSum %= 10

        cur.Val = digitSum
    }

    return dummy.Next
}
*/

func addRecursive(prev, l1, l2 *ListNode, carry int) {
    if l1 == nil && l2 == nil && carry == 0 {
        return
    }

    digitSum := carry

    if l1 != nil {
        digitSum += l1.Val
        l1 = l1.Next
    }

    if l2 != nil {
        digitSum += l2.Val
        l2 = l2.Next
    }

    carry = digitSum / 10
    digitSum %= 10

    cur := &ListNode{
        Val: digitSum,
    }
    prev.Next = cur

    addRecursive(cur, l1, l2, carry)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    addRecursive(dummy, l1, l2, 0)
    return dummy.Next
}