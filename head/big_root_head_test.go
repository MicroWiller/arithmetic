package head

import (
	"reflect"
	"testing"
)

func TestBigRootHeap(t *testing.T) {
	h := NewBigRootHeap()
	for i := 1; i <= 20; i++ {
		h.Push(Element(i))
	}
	for i := 20; i < 1; i++ {
		x := h.Pop()
		if x != Element(i) {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestGetLeastNumbers(t *testing.T) {
	arr := []int{3, 2, 1}
	k := 2
	expect1 := []int{1, 2}
	expect2 := []int{2, 1}

	result := GetLeastNumbers(arr, k)
	if !reflect.DeepEqual(result, expect1) {
		if !reflect.DeepEqual(result, expect2) {
			t.Errorf("expect %v or %v, but get %v", expect1, expect2, result)
		}
	}

}
