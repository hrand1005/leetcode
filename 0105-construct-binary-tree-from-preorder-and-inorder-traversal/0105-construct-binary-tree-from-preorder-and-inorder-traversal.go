/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    
    root := &TreeNode{
        Val: preorder[0],
    }
    
    rootIndex := 0
    for inorder[rootIndex] != root.Val {
        rootIndex++
    }
    
    inorderLeft := inorder[:rootIndex]
    inorderRight := inorder[rootIndex+1:]
    
    root.Left = buildTree(preorder[1:len(inorderLeft)+1], inorderLeft)
    root.Right = buildTree(preorder[len(inorderLeft)+1:], inorderRight)
    
    return root
}