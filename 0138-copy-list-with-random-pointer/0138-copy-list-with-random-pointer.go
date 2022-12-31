/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    
    headCopy := &Node{
        Val: head.Val,
    }
    cur, curCopy := head, headCopy
    origToCopy := make(map[*Node]*Node)
    
    for cur != nil {
        if cur.Next != nil {
            curCopy.Next = &Node{
                Val: cur.Next.Val,
            }
        }
        origToCopy[cur] = curCopy
        cur = cur.Next
        curCopy = curCopy.Next
    }
    
    cur, curCopy = head, headCopy
    for cur != nil {
        if cur.Random != nil {
            curCopy.Random = origToCopy[cur.Random]
        }
        cur = cur.Next
        curCopy = curCopy.Next
    }
    
    return headCopy
}