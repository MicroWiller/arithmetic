package list

// 206. 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
// 提示：
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Value: 0,
		Next:  nil,
	}
	for head != nil {
		tmp := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = tmp
	}
	return dummy.Next
}
