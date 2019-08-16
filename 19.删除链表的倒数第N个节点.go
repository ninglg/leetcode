/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    ha, hb := head, head
    // 先移动快指针
    for i := 0; i < n; i++ {
        hb = hb.Next
    }
    
    // 同时移动快慢指针
    for hb != nil && hb.Next != nil {
        ha = ha.Next
        hb = hb.Next
    }
    
    // 针对特定边界情况
    if ha == head && hb == nil {
        return ha.Next
    }
    
    ha.Next = ha.Next.Next
    
    return head
}
