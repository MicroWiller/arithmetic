package main

import (
	"sync"
	"testing"
)

func TestFishEat(t *testing.T) {
	size := [][]int{
		[]int{4, 2, 5, 3, 1},
		[]int{4, 3, 2, 1, 5},
	}
	dir := [][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{0, 1, 0, 0, 0},
	}
	ans := []int{3, 2}

	for i := 0; i < len(ans); i++ {
		getAns := FishEat(size[i], dir[i])
		if getAns != ans[i] {
			t.Errorf("expect answer is %d, actually get is %d", ans[i], getAns)
		}
	}
}

func TestIsValidWithBracket(t *testing.T) {
	s := "()()(())"
	getAns := IsValidWithBracket(s)
	if !getAns {
		t.Errorf("expect answer is true, actually get is %v", getAns)
	}
}

func TestIsValidWithCount(t *testing.T) {
	s := "()()(())"
	getAns := IsValidWithCount(s)
	if !getAns {
		t.Errorf("expect answer is true, actually get is %v", getAns)
	}
}

func TestIsValidWithKinds(t *testing.T) {
	s := "[{}({})]"
	getAns := IsValidWithKinds(s)
	if !getAns {
		t.Errorf("expect answer is true, actually get is %v", getAns)
	}
}


func TestStack_Push(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(10)
	stack := NewStack()
	for i := 0; i < 10; i++ {
		go func(i int) {
			stack.Push(i)
			group.Done()
		}(i)
	}
	group.Wait()
	if stack.len != 10 {
		t.Errorf("expect answer is %d, actually get is %d", 10, stack.len)
	}
}

