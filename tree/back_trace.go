package tree

/*
 * [113] 路径总和 II
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 * https://www.lintcode.com/problem/path-sum-ii/description
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 *              5
 *             / \
 *            4   8
 *           /   / \
 *          11  13  4
 *         /  \    / \
 *        7    2  5   1
 *
 * 返回:
 * [
 *   [5,4,11,2],
 *   [5,8,4,5]
 * ]
 */
var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var path []int
	backTrace(root, targetSum, 0, path)
	return res
}

func backTrace(node *TreeNode, targetSum int, sum int, path []int) {
	if node == nil {
		return
	}
	sum += node.Val
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil {
		if sum == targetSum {
			res = append(res, path)
		}
	} else {
		backTrace(node.Left, targetSum, sum, path)
		backTrace(node.Right, targetSum, sum, path)
	}
	path = append(path, len(path)-1)
}
