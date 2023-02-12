/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    diam := 0
    
    var maxDepth func(*TreeNode) int
    maxDepth = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        maxLeft := maxDepth(root.Left)
        maxRight := maxDepth(root.Right)
        diam = max(diam, maxLeft + maxRight)
        return max(maxLeft, maxRight) + 1
    }
    
    maxDepth(root)
    return diam
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}