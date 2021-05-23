package queue

import "container/list"

// 单调递减队列：
// 1. 入队，扩张单调队列的覆盖范围；
// 2. 出队，控制单调递减队列的覆盖范围；
// 3. 队首元素就是覆盖范围的最大值；
// 4. 队列中的元素个数小于覆盖范围元素的个数。

// MonotonicQueueReduce 非严格单调递减队列
type MonotonicQueueReduce struct {
	queue *list.List // 使用双端队列来实现
}

func NewMonotonicQueueReduce() *MonotonicQueueReduce {
	return &MonotonicQueueReduce{
		queue: list.New(),
	}
}

// Push 后进
func (m *MonotonicQueueReduce) Push(val int) {
	for m.queue.Len() != 0 && m.queue.Back().Value.(int) < val {
		m.queue.Remove(m.queue.Back())
	}
	m.queue.PushBack(val)
}

// Pop 移除队头，出队时候，相等才出队
func (m *MonotonicQueueReduce) Pop(val int) {
	if m.queue.Len() != 0 && m.queue.Front().Value.(int) == val {
		m.queue.Remove(m.queue.Front())
	}
}

// IsEmpty 判断队列是否为空
func (m *MonotonicQueueReduce) IsEmpty() bool {
	return m.queue.Len() == 0
}

// 例 3：滑动窗口的最大值
// 【题目】给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。
// 输入：nums = [1,3,-1,-3,5,3], k = 3
// 输出：[3,3,5,5]
// [239] 滑动窗口最大值：https://leetcode-cn.com/problems/sliding-window-maximum/description/
func MaxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return nil
	}
	m := NewMonotonicQueueReduce()
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		m.Push(nums[i])
		if i < k-1 {
			continue
		}
		res = append(res, m.queue.Front().Value.(int))
		m.Pop(nums[i-k+1])
	}
	return res
}

// 扩展： 用循环队列来模拟单调队列，求解滑动窗口最大值的题目

// 扩展：求滑动窗口最小值
//【题目】给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最小值。
// 输入：nums = [1,3,-1,-3,5,3], k = 3
// 输出：[-1,-3,-3,-3]
func MinSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	// 单调递增的队头，是最小值
	queue := NewMonotonicQueueIncrease()
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		queue.Push(nums[i])
		if i < k-1 {
			continue
		}
		res = append(res, queue.queue.Front().Value.(int))
		queue.Pop(nums[i-k+1])
	}
	return res
}

// 单调递增队列
// 队首元素, 是覆盖范围的最小值
type MonotonicQueueIncrease struct {
	queue *list.List
}

// Push 前进
func (m *MonotonicQueueIncrease) Push(val int) {
	if m.queue.Len() != 0 &&
		// 当 val <= 队列的头节点，加入队头
		val > m.queue.Front().Value.(int) {
		return
	}
	m.queue.PushFront(val)
}

// Pop 移除队尾
func (m *MonotonicQueueIncrease) Pop(val int) {
	if m.queue.Len() != 0 && m.queue.Back().Value.(int) == val {
		m.queue.Remove(m.queue.Back())
	}
}

func (m *MonotonicQueueIncrease) IsEmpty() bool {
	return m.queue.Len() == 0
}

func NewMonotonicQueueIncrease() *MonotonicQueueIncrease {
	return &MonotonicQueueIncrease{list.New()}
}

//  例 4：捡金币游戏
// 【题目】给定一个数组 A[]，每个位置 i 放置了金币 A[i]，
//  小明从 A[0] 出发。当小明走到 A[i] 的时候，下一步他可以选择 A[i+1, i+k]（当然，不能超出数组边界）。
//  每个位置一旦被选择，将会把那个位置的金币收走（如果为负数，就要交出金币）。
//  请问，最多能收集多少金币？
//  输入：[1,-1,-100,-1000,100,3], k = 2
//  输出：4
//  1696. 跳跃游戏 VI: https://leetcode-cn.com/problems/jump-game-vi/description/
func MaxResult(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k <= 0 {
		return 0
	}
	// 单调队列存储 i+k 个元素
	m := NewMonotonicQueueReduce()
	length := len(nums)
	// 存储当前下标下，收集到的最多金币
	get := make([]int, length)
	for i := 0; i < length; i++ {
		// 为当前下标腾位置，nums[i-k-1]不在当前下标选择范围内
		if i-k > 0 {
			m.Pop(get[i-k-1])
		}
		var old int
		if m.IsEmpty() {
			old = 0
		} else {
			old = m.queue.Front().Value.(int)
		}
		get[i] = old + nums[i]
		m.Push(get[i])
	}
	return get[length-1]
}
