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
    _, balanced := depthBalance(root)
    return balanced
}

func depthBalance(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    depthLeft, leftBalanced := depthBalance(root.Left)
    depthRight, rightBalanced := depthBalance(root.Right)
    if !leftBalanced || !rightBalanced || 1 < abs(depthLeft - depthRight) {
        return -1, false
    }
    return 1 + max(depthLeft, depthRight), true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
