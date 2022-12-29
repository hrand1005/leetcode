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

func zigzagLevelOrder(root *TreeNode) [][]int {
    queue := []*TreeNodeWithLevel{
        { root, 0 },
    }
    zigzag := [][]int{}
    
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.TreeNode == nil {
            continue
        }
        
        if len(zigzag) - 1 < node.Level {
            zigzag = append(zigzag, []int{node.Val})
        } else {
            if node.Level % 2 == 0 {
                zigzag[node.Level] = append(zigzag[node.Level], node.Val)
            } else {
                zigzag[node.Level] = append([]int{node.Val}, zigzag[node.Level]...)
            }
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
    
    return zigzag
}