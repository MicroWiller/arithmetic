package stack

import (
	"sync"
)

// 栈的总结：https://cdn.jsdelivr.net/gh/MicroWiller/photobed@master/StackWithGo.png

// Size[i] 表示第 i 条鱼的大小，Dir[i] 表示鱼的方向 （0 表示向左游，1 表示向右游）
// 鱼的行为符合以下几个条件:
// 1. 所有的鱼都同时开始游动，每次按照鱼的方向，都游动一个单位距离；
// 2. 当方向相对时，大鱼会吃掉小鱼；
// 3. 鱼的大小都不一样。
func FishEat(size, dir []int) int {
	if size == nil || dir == nil || len(size) != len(dir) {
		return -1
	}
	left, right := 0, 1
	stack := NewStack()
	totalFish := len(size)
	for i := 0; i < totalFish; i++ {
		curFishSize := size[i]
		curFishDir := dir[i]
		hasEat := false
		// 如果栈中还有鱼，并且栈中鱼向右，当前的鱼向左游，那么就会有相遇的可能性
		for !stack.Empty() &&
			dir[stack.Peak().(int)] == right &&
			curFishDir == left {
			// 当前🐟 > 栈中🐟 ，栈中🐟被吃掉
			if curFishSize > size[stack.Peak().(int)] {
				stack.Pop()
			} else {
				// 当前🐟 < 栈中🐟 ，当前🐟被吃掉
				hasEat = true
				break
			}
		}
		if !hasEat {
			stack.Push(i)
		}
	}
	return stack.len
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
// 有效字符串需满足：
// 1. 左括号必须用相同类型的右括号闭合
// 2. 左括号必须以正确的顺序闭合
// 3. 注意空字符串可被认为是有效字符串
func IsValidWithKinds(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack.Push(s[i])
		case ')':
			if stack.Empty() {
				return false
			}
			if stack.Peak().(uint8) != '(' {
				return false
			}
			stack.Pop()
		case '{':
			stack.Push(s[i])
		case '}':
			if stack.Empty() {
				return false
			}
			if stack.Peak().(uint8) != '{' {
				return false
			}
			stack.Pop()
		case '[':
			stack.Push(s[i])
		case ']':
			if stack.Empty() {
				return false
			}
			if stack.Peak().(uint8) != '[' {
				return false
			}
			stack.Pop()
		}
	}
	return true
}

// 判断字符串括号是否合法
func IsValidWithBracket(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := NewStack()
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		curChar := s[i]
		if curChar == '(' {
			stack.Push(curChar)
		} else if curChar == ')' {
			if stack.Empty() {
				return false
			}
			stack.Pop()
		}
	}
	return stack.Empty()
}

func IsValidWithCount(s string) bool {
	if len(s) == 0 {
		return false
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			if count < 0 {
				return false
			}
			count--
		}
	}
	return count == 0
}

type (
	Stack struct {
		top  *Node
		len  int
		lock *sync.RWMutex
	}
	Node struct {
		prev  *Node
		value interface{}
	}
)

func NewStack() *Stack {
	return &Stack{
		top:  nil,
		len:  0,
		lock: &sync.RWMutex{},
	}
}
func (s *Stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	newNode := &Node{s.top, value}
	s.top = newNode
	s.len++
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len == 0 {
		return nil
	}
	nowTop := s.top
	s.top = nowTop.prev
	s.len--
	return nowTop.value
}

func (s *Stack) Peak() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.len == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Empty() bool {
	return s.len == 0
}
