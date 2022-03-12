/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil {
        return nil
    }

    var ret []int

    for i := 0; i < len(root.Children); i++ {
        ret = append(ret, postorder(root.Children[i])...)
    }

    ret = append(ret, root.Val)

    return ret
}
