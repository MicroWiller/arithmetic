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

// [203] 移除链表元素：https://leetcode-cn.com/problems/remove-linked-list-elements/description/
//【题目】给定一个链表头及一个整数值，要求把链表里面等于整数值的结点都从链表中移除出去。
// 输入：1->2->3->2->4, remove = 2
// 输出：1->3->4。
// 解释：要移除的整数值是 2。那么移除之后，返回的结果应该是 1->3->4。
// 示例:
// 输入: 1->2->6->3->4->5->6, val = 6
// 输出: 1->2->3->4->5
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Value: 0,
		Next:  nil,
	}
	tail := dummy
	for head != nil {
		tmp := head.Next
		if head.Value != val {
			// 尾插
			tail.Next = head
			// or： tail = head
			tail = tail.Next
		}
		head = tmp
	}
	tail.Next = nil
	return dummy.Next
}

// [83] 删除排序链表中的重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
// 给定一个排序链表，删除重复出现的元素，使得每个元素只出现一次。
// 输入: 1->1->2->3->3
// 输出: 1->2->3
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Value: 0,
		Next:  nil,
	}
	tail := dummy
	for head != nil {
		tmp := head.Next
		if tail == dummy || tail.Value != head.Value {
			tail.Next = head
			tail = head
		}
		head = tmp
	}
	tail.Next = nil
	return dummy.Next
}
