/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }
    l := make([]int,0)
    for head != nil {
        l = append(l,head.Val)
        head = head.Next
    }
    n := len(l)
    k = k%n
    for i:=0;i<k;i++ {
        for j:=n-1;j>0;j-- {
            tmp := l[j]
            l[j] = l[j-1]
            l[j-1] = tmp
        }
    }
    ret := &ListNode{}
    h := ret
    for i:=0;i<n;i++ {
        ret.Next = &ListNode{l[i],nil}
        ret = ret.Next
    }
    return h.Next
}

