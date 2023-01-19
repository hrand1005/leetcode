/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
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
*/

func hasPathSum(root *TreeNode, targetSum int) bool {
    return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, currentSum, targetSum int) bool {
    if root == nil {
        return false
    }
    currentSum += root.Val
    if root.Left == nil && root.Right == nil {
        return currentSum == targetSum
    }
    return dfs(root.Left, currentSum, targetSum) || dfs(root.Right, currentSum, targetSum)
}