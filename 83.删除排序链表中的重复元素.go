/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    h, p := head, head
    
    for h != nil && h.Next != nil {
        p = h.Next
        for p != nil && p.Val == h.Val {
            p = p.Next
        }
        h.Next = p
        h = p
    }
    
    return head
}
