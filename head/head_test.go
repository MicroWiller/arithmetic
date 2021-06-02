package head

import (
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
