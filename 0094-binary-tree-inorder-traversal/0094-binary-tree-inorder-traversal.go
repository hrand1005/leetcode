/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    left := append(inorderTraversal(root.Left), root.Val)
    return append(left, inorderTraversal(root.Right)...)
}
*/

type VisitedNode struct {
    Node *TreeNode
    Visited bool
}

func inorderTraversal(root *TreeNode) []int {
    inorderPath := []int{}
    
    stack := []*VisitedNode{
        &VisitedNode{
            Node: root,
            Visited: false,
        },
    }
    
    for len(stack) > 0 {
        this := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if this.Node != nil {
            if this.Visited {
                inorderPath = append(inorderPath, this.Node.Val)
            } else {
                stack = append(stack, &VisitedNode{this.Node.Right, false})
                stack = append(stack, &VisitedNode{this.Node, true})
                stack = append(stack, &VisitedNode{this.Node.Left, false})
            }
        }
    }
    
    return inorderPath
}