package tree

import "arithmetic/stack"

// TreeNode 二叉树
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
}

func PreOrder(node *TreeNode, ans []int) {
	if node != nil {
		// 前序
		ans = append(ans, node.Val)
		PreOrder(node.Left, ans)
		PreOrder(node.Right, ans)
	}
}

func PreOrderByStack(node *TreeNode) {
	if node == nil {
		return
	}
	newStack := stack.NewStack()
	var ans []int
	newStack.Push(node)
	for !newStack.Empty() {
		curNode := newStack.Pop().(*TreeNode)
		ans = append(ans, curNode.Val)
		if curNode.Left != nil {
			newStack.Push(curNode.Left)
		}
		if curNode.Right != nil {
			newStack.Push(curNode.Right)
		}
	}
}

func InOrder(node *TreeNode, ans []int) {
	if node != nil {
		PreOrder(node.Left, ans)
		// 中序
		ans = append(ans, node.Val)
		PreOrder(node.Right, ans)
	}
}

func InOrderByStack(node *TreeNode, ans []int) {
	if node == nil {
		return
	}
	newStack := stack.NewStack()
	for node != nil || !newStack.Empty() {
		for node != nil {
			newStack.Push(node)
			node = node.Left
		}
		curNode := newStack.Pop().(*TreeNode)
		ans = append(ans, curNode.Val)
		// 不用判断 curNode.Right 是否为nil, 第一次for循环会进行校验
		node = curNode.Right
	}
}

// Morris解法, 空间复杂度：O(1), 没有使用stack
// https://github.com/lagoueduCol/Algorithm-Dryad/blob/main/06.Tree/94.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E4%B8%AD%E5%BA%8F%E9%81%8D%E5%8E%86.morris.java?fileGuid=xxQTRXtVcqtHK6j8

func LastOrder(node *TreeNode, ans []int) {
	if node != nil {
		PreOrder(node.Left, ans)
		PreOrder(node.Right, ans)
		// 后序
		ans = append(ans, node.Val)
	}
}
