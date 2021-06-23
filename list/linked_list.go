package list

// Node 单链表节点
type Node struct {
	Value interface{}
	Next  *Node
}

/*
 * [707] 设计链表
 * - https://leetcode-cn.com/problems/design-linked-list/
 */
type MyLinkedList struct {
	dummy, tail *Node
	len         int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	dummyNode := &Node{
		Value: nil,
		Next:  nil,
	}
	return MyLinkedList{
		dummy: dummyNode,
		tail:  dummyNode,
		len:   0,
	}
}

// 获取当前下标的前一个节点
func (this *MyLinkedList) getPreNode(index int) *Node {
	front := this.dummy.Next
	back := this.dummy
	// 移动index-1次，从0开始计算；front可以移动到target
	for i := 0; i < index && front != nil; i++ {
		back = front
		front = front.Next
	}
	return back
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.len {
		return -1
	}
	preNode := this.getPreNode(index)
	return preNode.Next.Value.(int)
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		Value: val,
		Next:  nil,
	}
	newNode.Next = this.dummy.Next
	this.dummy.Next = newNode
	this.len++
	// 边界：考虑尾节点
	if this.dummy == this.tail {
		this.tail = newNode
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		Value: val,
		Next:  nil,
	}
	this.tail.Next = newNode
	this.tail = this.tail.Next
	this.len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	} else if index == this.len {
		this.AddAtTail(val)
	} else if index <= 0 {
		this.AddAtHead(val)
	} else {
		newNode := &Node{
			Value: val,
			Next:  nil,
		}
		preNode := this.getPreNode(index)
		newNode.Next = preNode.Next
		preNode.Next = newNode
		this.len++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}
	preNode := this.getPreNode(index)
	if this.tail == preNode.Next {
		this.tail = preNode
	}
	this.len--
	preNode.Next = preNode.Next.Next
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
