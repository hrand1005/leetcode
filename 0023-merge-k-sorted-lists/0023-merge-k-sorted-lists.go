/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
import (
    "math"
)

const notFound = -1

func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for {
        idx := minIdx(lists)
        if idx == notFound {
            break
        }
        cur.Next = lists[idx]
        lists[idx] = lists[idx].Next
        cur = cur.Next
    }
    return dummy.Next
}

func minIdx(lists []*ListNode) int {
    minVal := math.MaxInt32
    idx := notFound
    for i, n := range lists {
        if n != nil && n.Val < minVal {
            minVal = n.Val
            idx = i
        }
    }
    return idx
}
*/

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    midpoint := len(lists) / 2
    first := mergeKLists(lists[:midpoint])
    second := mergeKLists(lists[midpoint:])
    return mergeTwo(first, second)
}

func mergeTwo(first, second *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for first != nil && second != nil {
        if first.Val < second.Val {
            cur.Next = first
            first = first.Next
        } else {
            cur.Next = second
            second = second.Next
        }
        cur = cur.Next
    }
    if first == nil {
        cur.Next = second
    } else {
        cur.Next = first
    }
    return dummy.Next
}