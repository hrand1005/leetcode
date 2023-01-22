/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    preorder := []int{root.Val}
    left := preorderTraversal(root.Left)
    right := preorderTraversal(root.Right)
    return append(preorder, append(left, right...)...)
}
*/

func preorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{root}
    preorder := []int{}
    for len(stack) != 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node == nil {
            continue
        }
        preorder = append(preorder, node.Val)
        stack = append(stack, node.Right)
        stack = append(stack, node.Left)
    }
    return preorder
}