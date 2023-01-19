/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type NodeWithSum struct {
    *TreeNode
    Sum int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    stack := []*NodeWithSum{
        {
            TreeNode: root,
            Sum: root.Val,
        },
    }
    
    for len(stack) != 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if node.Left == nil && node.Right == nil && node.Sum == targetSum{
            return true
        }
        if node.Left != nil {
            stack = append(stack, &NodeWithSum{
                TreeNode: node.Left,
                Sum: node.Sum + node.Left.Val,
            })
        }
        if node.Right != nil {
            stack = append(stack, &NodeWithSum{
                TreeNode: node.Right,
                Sum: node.Sum + node.Right.Val,
            })
        }
    }
    
    return false
}