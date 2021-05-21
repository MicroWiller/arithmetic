package queue

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3}
	k := 3
	expect := []int{3, 3, 5, 5}
	res := MaxSlidingWindow(nums, k)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Expect: %v, but actually get: %v", expect, res)
	}
}
