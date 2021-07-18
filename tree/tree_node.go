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
	stack := stack.NewStack()
	var ans []int
	stack.Push(node)
	for !stack.Empty() {
		curNode := stack.Pop().(*TreeNode)
		ans = append(ans, curNode.Val)
		if curNode.Left != nil {
			stack.Push(curNode.Left)
		}
		if curNode.Right != nil {
			stack.Push(curNode.Right)
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

func LastOrder(node *TreeNode, ans []int) {
	if node != nil {
		PreOrder(node.Left, ans)
		PreOrder(node.Right, ans)
		// 后序
		ans = append(ans, node.Val)
	}
}
