/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{root}
    ret := []float64{float64(root.Val)}
    for len(queue) > 0 {
        l := len(queue)

        var sum,c int
        for l > 0 {
            if queue[0].Left != nil {
                queue = append(queue, queue[0].Left)
                sum += queue[0].Left.Val
                c += 1
            }

            if queue[0].Right != nil {
                queue = append(queue, queue[0].Right)
                sum += queue[0].Right.Val
                c += 1
            }

            l--
            queue = queue[1:]
        }

        if c > 0 {
            ret = append(ret, float64(sum)/float64(c))
        }
    }

    return ret
}
