/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    var ret []int
    ret = append(ret, root.Val)

    ret = append(ret, preorderTraversal(root.Left)...)
    ret = append(ret, preorderTraversal(root.Right)...)

    return ret
}

