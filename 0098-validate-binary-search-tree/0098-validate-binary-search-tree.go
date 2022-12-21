/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
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
*/

func isValidBST(root *TreeNode) bool {
    visitedOrder := inorderTraversal(root)
    
    for i := 1; i < len(visitedOrder); i++ {
        if visitedOrder[i] <= visitedOrder[i-1] {
            return false
        }
    }
    
    return true
}

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    leftVisited := inorderTraversal(root.Left)
    visited := append(leftVisited, root.Val)
    rightVisited := inorderTraversal(root.Right)
    
    return append(visited, rightVisited...)
}