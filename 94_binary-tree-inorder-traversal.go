/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    
    res := make([]int,0,10)
    var helper func(r *TreeNode)
    
    helper = func(r *TreeNode) {
        if r == nil{
            return 
        }
        helper(r.Left)
        res = append(res,r.Val)
        helper(r.Right)
    }
    
    helper(root)
    return res
}

