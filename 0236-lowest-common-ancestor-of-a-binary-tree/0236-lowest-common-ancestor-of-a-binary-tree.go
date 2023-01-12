/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root == p || root == q {
        return root
    }

    var left *TreeNode
    if root.Left != nil {
        left = lowestCommonAncestor(root.Left, p, q)
    }

    var right *TreeNode
    if root.Right != nil {
        right = lowestCommonAncestor(root.Right, p, q)
    }

    if left != nil && right != nil {
        return root
    }
    
    if left != nil {
        return left
    }
    return right
}