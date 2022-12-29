/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
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
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
    queue := []*TreeNode{ root }
    level := 1
    zigzag := [][]int{}
    
    for len(queue) != 0 {
        levelCount := len(queue)
        vals := make([]int, 0, levelCount)
        newQueue := make([]*TreeNode, 0, levelCount*2)
        
        for i := 0; i < levelCount; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node == nil {
                continue
            }
            
            vals = append(vals, node.Val)
            
            if level % 2 == 0 {
                newQueue = append([]*TreeNode{node.Right}, newQueue...)
                newQueue = append([]*TreeNode{node.Left}, newQueue...)
            } else {
                newQueue = append([]*TreeNode{node.Left}, newQueue...)
                newQueue = append([]*TreeNode{node.Right}, newQueue...)
            }
        }
        
        if len(vals) != 0 {
            zigzag = append(zigzag, vals)
        }
        queue = newQueue
        level++
    }
    
    return zigzag
}