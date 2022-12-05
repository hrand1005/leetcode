/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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