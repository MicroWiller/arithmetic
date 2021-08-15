package tree

import (
	"math"
)

/*
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 * https://www.lintcode.com/problem/validate-binary-search-tree/
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 假设一个二叉搜索树具有如下特征：
 *
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 * 示例 1:
 * 输入:
 *    2
 *   / \
 *  1   3
 * 输出: true
 *
 * 示例 2:
 * 输入:
 *    5
 *   / \
 *  1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
 */
var isBST = true

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

func IsValidBST(root *TreeNode) bool {
	isBST = true
	isValidBstByPre(root, IntMin, IntMax)
	return isBST
}
func isValidBstByPre(root *TreeNode, l, r int) {
	if root == nil || !isBST {
		return
	}
	if !(l < root.Val && root.Val < r) {
		isBST = false
		return
	}
	isValidBstByPre(root.Left, l, root.Val)
	isValidBstByPre(root.Right, root.Val, r)
}

// IsValidBstIn 中序校验是不是二叉搜索树
func IsValidBstIn(root *TreeNode) bool {
	isBST = true
	preValue := IntMin
	isValidBstByIn(root, preValue)
	return isBST
}

func isValidBstByIn(node *TreeNode, preValue int) {
	if node == nil || !isBST {
		return
	}
	isValidBstByIn(node.Left, preValue)
	if preValue >= node.Val {
		isBST = false
		return
	}
	preValue = node.Val
	isValidBstByIn(node.Right, preValue)
}

/*
 * [100] 相同的树
 * https://leetcode-cn.com/problems/same-tree/description/
 * https://www.lintcode.com/problem/same-tree/
 *
 * 给定两个二叉树，编写一个函数来检验它们是否相同。
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 * 示例 1:
 * 输入:      1         1
 *          / \       / \
 *         2   3     2   3
 *
 *        [1,2,3],   [1,2,3]
 * 输出: true
 *
 * 示例 2:
 * 输入:      1         1
 *          /           \
 *         2             2
 *
 *       [1,2],     [1,null,2]
 * 输出: false
 *
 * 示例 3:
 * 输入:      1         1
 *          / \       / \
 *         2   1     1   2
 *
 *        [1,2,1],   [1,1,2]
 *
 * 输出: false
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

/*
 *
 * [572] 另一个树的子树
 * https://leetcode-cn.com/problems/subtree-of-another-tree/description/
 * https://www.lintcode.com/problem/subtree-of-another-tree/
 *
 * 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。
 * s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s
 * 也可以看做它自身的一棵子树。
 *
 * 示例 1:
 * 给定的树 s:
 *     3
 *    / \
 *   4   5
 *  / \
 * 1   2
 *
 * 给定的树 t：
 *   4
 *  / \
 * 1   2
 * 返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。
 *
 * 示例 2:
 * 给定的树 s：
 *     3
 *    / \
 *   4   5
 *  / \
 * 1   2
 *    /
 *   0
 *
 * 给定的树 t：
 *   4
 *  / \
 * 1   2
 * 返回 false。
 */
func isSubtree(s, t *TreeNode) bool {
	if s == t {
		return true
	}
	if s != nil && t == nil {
		return true
	}
	if s == nil {
		return false
	}

	return s.Val == t.Val && isSame(s, t) ||
		isSame(s.Left, t) || isSame(s.Right, t)
}

func isSame(a *TreeNode, b *TreeNode) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val &&
		isSame(a.Left, b.Left) &&
		isSame(a.Right, b.Right)
}

/*
 * [501] 二叉搜索树中的众数
 * https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/description/
 *
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 *
 * 假定 BST 有如下定义：
 * 结点左子树中所含结点的值小于等于当前结点的值
 * 结点右子树中所含结点的值大于等于当前结点的值
 * 左子树和右子树都是二叉搜索树
 *
 *
 * 例如：
 * 给定 BST [1,null,2,2],
 *
 *   1
 *    \
 *     2
 *    /
 *   2
 * 返回[2].
 * 提示：如果众数超过1个，不需考虑输出顺序
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	// 保存最大频类的节点
	maxFreq := IntMin
	// 前一个元素的值
	var preValue int
	// 前一个元素的计数
	var preCnt int

	midOrder4FindMode(root, res, maxFreq, preValue, preCnt)

	// ？？不是求最大的频类么？？
	return res[len(res)-1:]
}

func midOrder4FindMode(node *TreeNode, res []int, maxFreq int, preValue int, preCnt int) {
	if node == nil {
		return
	}
	midOrder4FindMode(node.Left, res, maxFreq, preValue, preCnt)
	if node.Val == preValue {
		preCnt++
	} else {
		preValue = node.Val
		preCnt = 1
	}
	if preCnt == maxFreq {
		res = append(res, node.Val)
	}
	maxFreq = int(math.Max(float64(preCnt), float64(maxFreq)))
	midOrder4FindMode(node.Right, res, maxFreq, preValue, preCnt)
}

/*
	783. 二叉搜索树节点最小距离：
	https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/

	给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
*/
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDst = IntMax
	preVal = IntMin >> 1
	findMinDistanceByMidOrder(root)
	return minDst
}

var preVal, minDst int

func findMinDistanceByMidOrder(node *TreeNode) {
	if node == nil {
		return
	}
	findMinDistanceByMidOrder(node.Left)
	dst := node.Val - preVal
	if dst < minDst {
		minDst = dst
	}
	preVal = node.Val
	findMinDistanceByMidOrder(node.Right)
}

/*
 * [99] 恢复二叉搜索树
 * https://leetcode-cn.com/problems/recover-binary-search-tree/description/
 *
 * 给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。
 * 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？
 *
 * 示例 1：
 * 输入：root = [1,3,null,null,2]
 * 输出：[3,1,null,null,2]
 * 解释：3 不能是 1 左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
 *
 * 示例 2：
 * 输入：root = [3,1,4,null,null,2]
 * 输出：[2,1,4,null,null,3]
 * 解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
 *
 * 提示：
 * 树上节点的数目在范围 [2, 1000] 内
 * -2^31
 */

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	first, second, pre = nil, nil, nil
	midOrder4RecoverTree(root)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

var first, second, pre *TreeNode

func midOrder4RecoverTree(node *TreeNode) {
	if node == nil {
		return
	}
	midOrder4RecoverTree(node.Left)
	if pre != nil && pre.Val > node.Val {
		if first == nil {
			first = pre
		}
		second = node
	}
	pre = node
	midOrder4RecoverTree(node.Right)
}

/*
 * [450] 删除二叉搜索树中的节点
 * https://leetcode-cn.com/problems/delete-node-in-a-bst/description/
 * https://www.lintcode.com/zh-cn/problem/remove-node-in-binary-search-tree/
 *
 * 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key
 * 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
 *
 * 一般来说，删除节点可分为两个步骤：
 * 首先找到需要删除的节点；
 * 如果找到了，删除它。
 *
 * 说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
 *
 * 示例:
 * root = [5,3,6,2,4,null,7]
 * key = 3
 *
 *    5
 *   / \
 *  3   6
 * / \   \
 *2   4   7
 *
 * 给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
 * 一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
 *
 *    5
 *   / \
 *  4   6
 * /     \
 *2       7
 * 另一个正确答案是 [5,2,6,null,4,null,7]。
 *
 *    5
 *   / \
 *  2   6
 *   \   \
 *    4   7
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// 找到待删除节点
		// 判断待删除节点，为叶子节点，还是非叶子节点
		if root.Left == nil && root.Right == nil { // 叶子节点，直接返回nil，使其上一层的Left/Right为nil
			return nil
		} else if root.Left != nil { // 非叶子节点，判断左右哪个树枝不为null
			// 找到左边最大值，交换，删除
			leftMax := root.Left
			for leftMax.Right != nil {
				leftMax = leftMax.Right
			}
			root.Val, leftMax.Val = leftMax.Val, root.Val
			root.Left = deleteNode(root.Left, key)
		} else if root.Right != nil {
			// 找到右边最小，交换，删除
			rightMin := root.Right
			for rightMin.Left != nil {
				rightMin = rightMin.Left
			}
			root.Val, rightMin.Val = rightMin.Val, root.Val
			root.Right = deleteNode(root.Right, key)
		}
	}
	return root
}
