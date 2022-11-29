/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    medianIndex := len(nums) / 2
    rootNode := &TreeNode{
        Val: nums[medianIndex],
    }
    left := sortedArrayToBST(nums[:medianIndex])
    right := sortedArrayToBST(nums[medianIndex+1:])
    
    rootNode.Left = left
    rootNode.Right = right
    
    return rootNode
}