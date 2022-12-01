/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    visited := map[*ListNode]int{}
    
    cur1 := headA
    cur2 := headB
    for cur1 != nil || cur2 != nil {
        if cur1 != nil {
            if _, ok := visited[cur1]; ok {
                return cur1
            } 
            visited[cur1]++
            cur1 = cur1.Next
        }
        
        if cur2 != nil {
            if _, ok := visited[cur2]; ok {
                return cur2
            }
            visited[cur2]++
            cur2 = cur2.Next
        }
    }
    
    return nil
}