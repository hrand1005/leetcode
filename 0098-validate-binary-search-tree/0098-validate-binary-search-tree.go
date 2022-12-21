/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValidBSTRecursive(root, nil, nil)
}

func isValidBSTRecursive(root *TreeNode, min, max *int) bool {
    if root == nil {
        return true
    }
    
    if min != nil && root.Val <= *min {
        return false
    }
    
    if max != nil && *max <= root.Val {
        return false
    }
    
    leftValid := isValidBSTRecursive(root.Left, min, &root.Val)
    rightValid := isValidBSTRecursive(root.Right, &root.Val, max)
    
    return leftValid && rightValid
}