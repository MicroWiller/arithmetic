package tree

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

var isBST = true

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

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
