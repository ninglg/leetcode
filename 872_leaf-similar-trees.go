/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    d1 := DFS(root1)
    d2 := DFS(root2)

    return reflect.DeepEqual(d1, d2)
}

func DFS(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    var ret []int

    if root.Left == nil && root.Right == nil {
        ret = append(ret, root.Val)
    }

    ret = append(ret, DFS(root.Left)...)
    ret = append(ret, DFS(root.Right)...)

    return ret
}
