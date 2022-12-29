/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type NodeWithLevel struct {
    *Node
    Level int
}

func connect(root *Node) *Node {
    queue := []*NodeWithLevel{
        {root, 0},
    }
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.Node == nil {
            continue
        }
        
        if len(queue) > 0 && queue[0].Level == node.Level {
            node.Next = queue[0].Node
        }
        
        queue = append(queue, &NodeWithLevel{
            node.Left,
            node.Level + 1,
        })
        queue = append(queue, &NodeWithLevel{
            node.Right,
            node.Level + 1,
        })
    }
    
    return root
}