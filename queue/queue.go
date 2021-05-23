package queue

// 总结：![](https://cdn.jsdelivr.net/gh/MicroWiller/photobed@master/QueueSummary.png)

// 二叉树
type TreeNode struct {
	value       int
	left, right *TreeNode
}

// 题目：从上到下按层打印二叉树，同一层结点按从左到右的顺序打印，每一层打印到一行。
func PrintBinaryTree(head *TreeNode) [][]int {
	if head == nil {
		return nil
	}
	// todo：使用切片实现队列
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	// check：添加是否添加到切片最后
	queue = append(queue, head)
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			// check：是否移除了第一个元素
			queue = append(queue[:0], queue[1:]...)
			level = append(level, curNode.value)
			if curNode.left != nil {
				queue = append(queue, curNode.left)
			}
			if curNode.right != nil {
				queue = append(queue, curNode.right)
			}
		}
		res = append(res, level)
	}
	return res
}

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
type Node struct {
	value             int
	left, right, next *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			// 移除头节点
			queue = append(queue[:0], queue[1:]...)
			if i < size-1 {
				curNode.next = queue[0]
			}
			if curNode.left != nil {
				queue = append(queue, curNode.left)
			}
			if curNode.right != nil {
				queue = append(queue, curNode.right)
			}
		}
	}
	return root
}

// 429. N 叉树的层序遍历
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	return nil
}
