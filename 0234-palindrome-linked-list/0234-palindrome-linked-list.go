/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
func isPalindrome(head *ListNode) bool {
    candidate := []int{}
    cur := head
    
    for cur != nil {
        candidate = append(candidate, cur.Val)
        cur = cur.Next
    }
    
    for i := 0; i < len(candidate)/2; i++ {
        if candidate[i] != candidate[len(candidate)-1-i] {
            return false
        }
    }
    
    return true
}
*/

func isPalindrome(head *ListNode) bool {
    var prev *ListNode
    slow := head
    fast := slow
    
    for fast != nil && fast.Next != nil {
        // advance 'fast' at double the rate
        fast = fast.Next
        fast = fast.Next
        
        // reverse the first half of the list while advancing
        slowNext := slow.Next
        slow.Next = prev
        prev = slow
        slow = slowNext
    }
    
    // advance past midpoint for odd-length lists
    if fast != nil {
        slow = slow.Next
    }
    
    firstHalf := prev
    secondHalf := slow
    
    for firstHalf != nil {
        if firstHalf.Val != secondHalf.Val {
            return false
        }
        
        firstHalf = firstHalf.Next
        secondHalf = secondHalf.Next
    }
    
    return true
}