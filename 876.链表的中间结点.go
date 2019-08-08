/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
    
	fast := head
	slow := head
    
	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
    
	return slow
}

