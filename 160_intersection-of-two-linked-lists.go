// 通过判断链表尾结点是否相同，即可知道是否相交
// 通过链表长度差，然后共同前进，判断在哪个节点最先相交
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    ha := headA
    hb := headB

    na := 1
    nb := 1

    for ha.Next != nil {
        ha = ha.Next
        na++
    }

    for hb.Next != nil {
        hb = hb.Next
        nb++
    }

    if ha != hb {
        return nil
    }

    ha = headA
    hb = headB

    if na <= nb {
        for i := 0; i < nb-na; i++ {
            hb = hb.Next
        }

        for i := 0; i < na; i++ {
            if ha == hb {
                return ha
            }
            ha = ha.Next
            hb = hb.Next
        }
    } else {
        for i := 0; i < na-nb; i++ {
            ha = ha.Next
        }

        for i := 0; i < nb; i++ {
            if ha == hb {
                return ha
            }
            ha = ha.Next
            hb = hb.Next
        }
    }

    return nil
}
