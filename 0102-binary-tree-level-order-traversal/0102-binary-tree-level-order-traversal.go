/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNodeWithLevel struct {
    *TreeNode
    Level int
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
    queue := []*TreeNodeWithLevel{
        { root, 0 },
    }
    values := [][]int{}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.TreeNode == nil {
            continue
        }
        
        if len(values) < node.Level+1 {
            values = append(values, []int{node.Val})
        } else {
            values[node.Level] = append(values[node.Level], node.Val)
        }
        
        queue = append(queue, &TreeNodeWithLevel{
            node.Left,
            node.Level + 1,
        })
        queue = append(queue, &TreeNodeWithLevel{
            node.Right,
            node.Level + 1,
        })
    }
    
    return values
}