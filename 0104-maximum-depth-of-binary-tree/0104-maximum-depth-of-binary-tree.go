/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func maxInt(ints []int) int {
    max := 0
    for _, v := range ints {
        if v > max {
            max = v
        }
    }
    return max
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    pathLengths := []int{
        maxDepth(root.Left) + 1,
        maxDepth(root.Right) + 1,
    }
    
    return maxInt(pathLengths)
}

/*
type DepthNode struct {
    Node *TreeNode
    Depth int
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    pathLengths := []int{}
    stack := []*DepthNode{
        {
            Node: root,
            Depth: 1,
        },
    }
    
    for len(stack) > 0 {
        this := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if this.Node.Left == nil && this.Node.Right == nil {
            pathLengths = append(pathLengths, this.Depth)
            continue
        }
        
        if this.Node.Left != nil {
            left := &DepthNode{
                Node: this.Node.Left, 
                Depth: this.Depth+1,
            }
            stack = append(stack,  left)
        }
        
        if this.Node.Right != nil {
            right := &DepthNode{
                Node: this.Node.Right, 
                Depth: this.Depth+1,
            }
            stack = append(stack,  right)
        }
    }
    
    return maxInt(pathLengths)
}
*/