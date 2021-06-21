package head

/*
 * [743] 网络延迟时间
 * https://leetcode-cn.com/problems/network-delay-time/description/
 *
 * 有 n 个网络节点，标记为 1 到 n。
 * 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
 * wi 是一个信号从源节点传递到目标节点的时间。
 * 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
 *
 * 示例 1：
 * 输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
 * 输出：2
 *
 * 示例 2：
 * 输入：times = [[1,2,1]], n = 2, k = 1
 * 输出：1
 *
 * 示例 3：
 * 输入：times = [[1,2,1]], n = 2, k = 2
 * 输出：-1
 *
 * 提示：
 * 1
 * 1
 * times[i].length == 3
 * 1 i, vi
 * ui != vi
 * 0 i
 * 所有 (ui, vi) 对都 互不相同（即，不含重复边）
 */

const INT_MAX = int(^uint(0) >> 1)

type edge struct {
	to   int
	cost int
}

func networkDelayTime(times [][]int, n, k int) int {
	if times == nil || len(times) == 0 {
		return -1
	}
	// 封装数据源
	allEdge := make([][]edge, 0, len(times))
	for _, time := range times {
		index := time[0]
		to := time[1]
		cost := time[2]
		allEdge[index] = append(allEdge[index], edge{to: to, cost: cost})
	}
	return 0
}
