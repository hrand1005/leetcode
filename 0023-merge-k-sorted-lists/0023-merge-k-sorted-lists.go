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

func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for {
        if idx, ok := minIdx(lists); ok {
            cur.Next = lists[idx]
            lists[idx] = lists[idx].Next
            cur = cur.Next
        } else {
            break
        }
    }
    return dummy.Next
}

func minIdx(lists []*ListNode) (int, bool) {
    minVal := math.MaxInt32
    idx, ok := 0, false
    
    for i, n := range lists {
        if n != nil && n.Val < minVal {
            ok = true
            minVal = n.Val
            idx = i
        }
    }
    return idx, ok
}