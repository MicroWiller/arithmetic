package list

// ListNode 单链表节点
type ListNode struct {
	Value int
	Next  *ListNode
}

/*
 * [707] 设计链表
 * - https://leetcode-cn.com/problems/design-linked-list/
 */
type MyLinkedList struct {
	dummy, tail *ListNode
	len         int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	dummyListNode := &ListNode{
		Value: 0,
		Next:  nil,
	}
	return MyLinkedList{
		dummy: dummyListNode,
		tail:  dummyListNode,
		len:   0,
	}
}

// 获取当前下标的前一个节点
func (this *MyLinkedList) getPreListNode(index int) *ListNode {
	front := this.dummy.Next
	back := this.dummy
	// 移动index-1次，从0开始计算；front可以移动到target
	for i := 0; i < index && front != nil; i++ {
		back = front
		front = front.Next
	}
	return back
}

/** Get the value of the index-th ListNode in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.len {
		return -1
	}
	preListNode := this.getPreListNode(index)
	return preListNode.Next.Value
}

/** Add a ListNode of value val before the first element of the linked list. After the insertion, the new ListNode will be the first ListNode of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newListNode := &ListNode{
		Value: val,
		Next:  nil,
	}
	newListNode.Next = this.dummy.Next
	this.dummy.Next = newListNode
	this.len++
	// 边界：考虑尾节点
	if this.dummy == this.tail {
		this.tail = newListNode
	}
}

/** Append a ListNode of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newListNode := &ListNode{
		Value: val,
		Next:  nil,
	}
	this.tail.Next = newListNode
	this.tail = this.tail.Next
	this.len++
}

/** Add a ListNode of value val before the index-th ListNode in the linked list. If index equals to the length of linked list, the ListNode will be appended to the end of linked list. If index is greater than the length, the ListNode will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	} else if index == this.len {
		this.AddAtTail(val)
	} else if index <= 0 {
		this.AddAtHead(val)
	} else {
		newListNode := &ListNode{
			Value: val,
			Next:  nil,
		}
		preListNode := this.getPreListNode(index)
		newListNode.Next = preListNode.Next
		preListNode.Next = newListNode
		this.len++
	}
}

/** Delete the index-th ListNode in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}
	preListNode := this.getPreListNode(index)
	if this.tail == preListNode.Next {
		this.tail = preListNode
	}
	this.len--
	preListNode.Next = preListNode.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
