package head

import (
	"testing"
)

func TestFurthestBuilding(t *testing.T) {
	heights := []int{4, 2, 7, 6, 9, 14, 12}
	bricks := 5
	ladders := 1
	expect := 4
	ans := FurthestBuilding(heights, bricks, ladders)
	if ans != expect {
		t.Errorf("expect %v , but get %v", expect, ans)
	}
}
