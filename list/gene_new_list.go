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

// [82] 删除排序链表中的重复元素 II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
//  给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//  示例 1:
//  输入: 1->2->3->3->4->4->5
//  输出: 1->2->5
//  示例 2:
//  输入: 1->1->1->2->3
//  输出: 2->3
func deleteDuplicates2(head *ListNode) *ListNode {
	if head != nil {
		return nil
	}
	tmp := &ListNode{
		Value: 0,
		Next:  nil,
	}
	tmp_tail := tmp
	ans := &ListNode{
		Value: 0,
		Next:  nil,
	}
	ans_tail := ans
	for head != nil {
		back := head.Next
		if tmp == tmp_tail || tmp_tail.Value == head.Value {
			tmp_tail.Next = head
			tmp_tail = head
		} else {
			// 有且仅有一个节点
			if tmp.Next == tmp_tail {
				ans_tail.Next = tmp_tail
				ans_tail = tmp_tail
			}
			// 多个节点时，需要清空
			tmp_tail = tmp
			tmp_tail.Next = nil

			tmp_tail.Next = head
			tmp_tail = head
		}
		head = back
	}
	if tmp.Next == tmp_tail {
		ans_tail.Next = tmp_tail
		ans_tail = tmp_tail
	}
	ans_tail.Next = nil
	return ans.Next
}

/*
 * [21] 合并两个有序链表
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例 1：
 * 输入：l1 = [1,2,4], l2 = [1,3,4]
 * 输出：[1,1,2,3,4,4]
 * 示例 2：
 * 输入：l1 = [], l2 = []
 * 输出：[]
 * 示例 3：
 * 输入：l1 = [], l2 = [0]
 * 输出：[0]
 */
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Value: 0,
		Next:  nil,
	}
	tail := dummy
	for l1 != nil || l2 != nil {
		if l2 == nil || l1 != nil && l1.Value < l2.Value {
			tail.Next = l1
			tail = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = l2
			l2 = l2.Next
		}
	}
	tail.Next = nil
	return dummy.Next
}
