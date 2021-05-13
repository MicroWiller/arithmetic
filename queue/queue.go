package queue

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
