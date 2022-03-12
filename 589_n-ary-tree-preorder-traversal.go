/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return nil
    }

    var ret []int
    ret = append(ret, root.Val)

    for i := 0; i < len(root.Children); i++ {
        ret = append(ret, preorder(root.Children[i])...)
    }

    return ret
}

