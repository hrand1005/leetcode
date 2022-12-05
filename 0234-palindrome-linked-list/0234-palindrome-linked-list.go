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
        fast = fast.Next
        fast = fast.Next
        
        slowNext := slow.Next
        slow.Next = prev
        prev = slow
        slow = slowNext
    }
    
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