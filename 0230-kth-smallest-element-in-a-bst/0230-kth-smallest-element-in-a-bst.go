/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    smallest := make([]int, 0, k)
    smallest = kthSmallestRecursive(root, k, smallest)
    return smallest[k-1]
}

func kthSmallestRecursive(root *TreeNode, k int, smallest []int) []int {
    if root == nil {
        return smallest
    }
    if len(smallest) < k {
        smallest = kthSmallestRecursive(root.Left, k, smallest)
    }
    if len(smallest) < k {
        smallest = append(smallest, root.Val)
    }
    if len(smallest) < k {
        smallest = kthSmallestRecursive(root.Right, k, smallest)
    }
    return smallest
}