/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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