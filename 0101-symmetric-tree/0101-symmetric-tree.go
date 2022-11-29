/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
func checkMirrored(n1, n2 *TreeNode) bool {
    if n1 == nil && n2 == nil {
        return true
    }
    
    if n1 == nil || n2 == nil {
        return false
    }
    
    if n1.Val != n2.Val {
        return false
    }
    
    return checkMirrored(n1.Left, n2.Right) && checkMirrored(n1.Right, n2.Left)
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return checkMirrored(root.Left, root.Right)
}
*/


func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    stack := []*TreeNode{root.Left, root.Right}
    
    for len(stack) > 0 {
        n1, n2 := stack[len(stack)-1], stack[len(stack)-2]
        stack = stack[:len(stack)-2] // 'pop' stack elements
        
        if n1 == nil && n2 == nil {
            continue
        }
        
        if n1 == nil || n2 == nil {
            return false
        }
        
        if n1.Val != n2.Val {
            return false
        }
        
        stack = append(stack, n1.Left, n2.Right, n1.Right, n2.Left) 
    }
    
    return true
}