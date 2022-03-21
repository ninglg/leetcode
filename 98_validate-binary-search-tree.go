/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil{
        return true
    }
    var stack []*TreeNode
    cur := root
    var prev *TreeNode
    
    for len(stack)>0 || cur!=nil{
        if cur!=nil{
            stack = append(stack,cur)
            cur = cur.Left
        }else{
            cur = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if prev!=nil &&prev.Val >= cur.Val{
                return false
            }
            
            prev = cur
            cur = cur.Right
        }
    }
    
    return true
}

