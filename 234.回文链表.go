/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    isPalindrome := true

    //链表长度
    length := 0
    for cur := head; cur != nil; cur = cur.Next {
        length++
    }

    //将前半部分反转
    step := length / 2
    var prev *ListNode
    cur := head
    for i := 1; i <= step; i++ {
        cur.Next, prev, cur = prev, cur, cur.Next
    }
    mid := cur

    var left, right *ListNode = prev, nil
    if length%2 == 0 {
        //长度为偶数
        right = mid
    } else {
        right = mid.Next
    }

    //从中间向左右两边遍历对比
    for left != nil && right != nil {
        if left.Val != right.Val {
            //值不相等,不是回文链表
            isPalindrome = false
            break
        }
        left = left.Next
        right = right.Next
    }

    //将前半部分反转的链表进行复原
    cur = prev
    prev = mid
    for cur != nil {
        cur.Next, prev, cur = prev, cur, cur.Next
    }

    return isPalindrome
}

