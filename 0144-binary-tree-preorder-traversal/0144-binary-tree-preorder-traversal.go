/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    preorder := []int{root.Val}
    left := preorderTraversal(root.Left)
    right := preorderTraversal(root.Right)
    return append(preorder, append(left, right...)...)
}