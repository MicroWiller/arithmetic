package head

import "testing"

func TestEatenApples(t *testing.T) {
	apples := []int{1, 2, 3, 5, 2}
	days := []int{3, 2, 1, 4, 2}
	expect := 7
	ans := EatenApples(apples, days)
	if expect != ans {
		t.Errorf("expect is %v, but get %v", expect, ans)
	}
}
