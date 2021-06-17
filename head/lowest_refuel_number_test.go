package head

import "testing"

func TestMinRefuelStops(t *testing.T) {
	target := 100
	startFuel := 10
	stations := [][]int{
		[]int{10, 60},
		[]int{20, 30},
		[]int{30, 30},
		[]int{60, 40},
	}
	expect := 2
	ans := MinRefuelStops(target, startFuel, stations)
	if expect != ans {
		t.Errorf("expect is %v, but get %v", expect, ans)
	}
}
