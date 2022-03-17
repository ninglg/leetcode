/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    max := 0
    for i := 0; i < len(root.Children); i++ {
        temp := maxDepth(root.Children[i])
        if  temp > max {
            max = temp
        }
    }

    return max + 1
}

