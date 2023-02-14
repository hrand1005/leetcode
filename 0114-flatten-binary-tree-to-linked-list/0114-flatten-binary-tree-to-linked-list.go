/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    flatten(root.Right)
    flatten(root.Left)
    
    temp := root.Right
    root.Right = root.Left
    root.Left = nil
    
    cur := root
    for cur.Right != nil {
        cur = cur.Right
    }
    cur.Right = temp
}