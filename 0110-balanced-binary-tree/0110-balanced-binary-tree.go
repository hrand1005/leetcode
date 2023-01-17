/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    depthDiff := depth(root.Left) - depth(root.Right)
    if depthDiff < -1 || 1 < depthDiff {
        return false
    }
    return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}